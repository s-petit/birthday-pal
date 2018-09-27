package email

import (
	"bytes"
	"github.com/s-petit/birthday-pal/remind"
	"strings"
	"text/template"
	"time"
)

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

func toMail(contact remind.ContactBirthday, language string) ([]byte, error) {

	var i18nTemplate i18nTemplate = enTemplate{}
	if strings.ToUpper(language) == "FR" {
		i18nTemplate = frTemplate{}
	}

	return resolveMail(contact, i18nTemplate)
}

func resolveMail(contact remind.ContactBirthday, i18nTemplate i18nTemplate) ([]byte, error) {
	subjFuncs := template.FuncMap{
		"yearValid": yearValid,
	}

	bodyFuncs := template.FuncMap{
		"formatDate": i18nTemplate.formatDate,
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

	var subject = resolveTemplate(subj, contact)
	var body = resolveTemplate(bod, contact)

	m := subjectBody{subject.String(), body.String()}

	resolvedMail := resolveTemplate(mail, m)
	return resolvedMail.Bytes(), nil
}

func resolveTemplate(template *template.Template, object interface{}) bytes.Buffer {
	var doc bytes.Buffer
	template.Execute(&doc, object)
	return doc
}
