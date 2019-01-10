package email

import (
	"bytes"
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

func toMail(emailContacts Contacts, language string) ([]byte, error) {

	var i18nTemplate i18nTemplate = enTemplate{}
	if strings.ToUpper(language) == "FR" {
		i18nTemplate = frTemplate{}
	}

	return resolveMail(emailContacts, i18nTemplate)
}

func resolveMail(emailContacts Contacts, i18nTemplate i18nTemplate) ([]byte, error) {
	subjFuncs := template.FuncMap{
		"formatDate": i18nTemplate.formatDate,
	}

	bodyFuncs := template.FuncMap{
		"yearValid":  yearValid,
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
