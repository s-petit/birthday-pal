package email

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

type mTemplate interface {
	subject() string
	body() string
	dateLayout() string
	formatDate(date time.Time) string
}

// MUST match RFC-822 format
const mailFormat = `To: Birthday Pals
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

	var mailTemplate mTemplate = English{}
	if strings.ToUpper(language) == "FR" {
		mailTemplate = French{}
	}

	return resolveMail(emailContacts, mailTemplate)
}

func resolveMail(emailContacts Contacts, mailTemplate mTemplate) ([]byte, error) {
	subjFuncs := template.FuncMap{
		"formatDate": mailTemplate.formatDate,
	}

	bodyFuncs := template.FuncMap{
		"yearValid":  yearValid,
		"formatDate": mailTemplate.formatDate,
	}

	subj, err := template.New("subject").Funcs(subjFuncs).Parse(mailTemplate.subject())
	if err != nil {
		return nil, err
	}
	bod, err := template.New("body").Funcs(bodyFuncs).Parse(mailTemplate.body())
	if err != nil {
		return nil, err
	}
	mail, err := template.New("mail").Parse(mailFormat)
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
