package i18n

import (
	"bytes"
	"github.com/s-petit/birthday-pal/app/contact"
	"github.com/s-petit/birthday-pal/app/email/i18n/en"
	"github.com/s-petit/birthday-pal/app/email/i18n/fr"
	"github.com/s-petit/birthday-pal/app/remind"
	"html/template"
	"strings"
	"time"
)


// Contacts holds every contacts related data necessary for the email content.
type Contacts struct {
	Contacts     []contact.Contact
	RemindParams remind.Criteria
}

type Template interface {
	subject() string
	body() string
	dateLayout() string
	formatDate(date time.Time) string
}

// MUST match RFC-822 format
const mailTemplate = `To: Birthday Pals
Subject: {{.Subject}}

{{.Body}}`

type subjectBody struct {
	Subject string
	Body    string
}

func yearValid(date time.Time) bool {
	return date.Year() > 0
}

func ToMail(emailContacts Contacts, language string) ([]byte, error) {

	var i18nTemplate Template = en.Template{}
	if strings.ToUpper(language) == "FR" {
		i18nTemplate = fr.Template{}
	}

	return resolveMail(emailContacts, i18nTemplate)
}

func resolveMail(emailContacts Contacts, i18nTemplate Template) ([]byte, error) {
	subjFuncs := template.FuncMap{
		"formatDate": Template.formatDate,
	}

	bodyFuncs := template.FuncMap{
		"yearValid":  yearValid,
		"formatDate": Template.formatDate,
	}

	subj, err := template.New("subject").Funcs(subjFuncs).Parse(i18nTemplate.subject())
	if err != nil {
		return nil, err
	}
	bod, err := template.New("body").Funcs(bodyFuncs).Parse(i18nTemplate.body())
	if err != nil {
		return nil, err
	}
	mail, err := template.New("mail").Parse(mailTemplate)
	if err != nil {
		return nil, err
	}

	var subject = resolveTemplate(subj, emailContacts)
	var body = resolveTemplate(bod, emailContacts)

	m := subjectBody{subject.String(), body.String()}

	resolvedMail := resolveTemplate(mail, m)
	return resolvedMail.Bytes(), nil
}

func resolveTemplate(template *template.Template, object interface{}) bytes.Buffer {
	var doc bytes.Buffer
	template.Execute(&doc, object)
	return doc
}
