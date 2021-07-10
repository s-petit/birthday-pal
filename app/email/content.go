package email

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

// MUST match RFC-5322 format
const mailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}`

type subjectBody struct {
	From    string
	To      string
	Subject string
	Body    string
}

func yearValid(date time.Time) bool {
	return date.Year() > 1900
}

func toMail(emailContacts Contacts, language string, sender string, recipients []string) ([]byte, error) {

	var i18nTemplate i18nTemplate = enTemplate{}
	if strings.ToUpper(language) == "FR" {
		i18nTemplate = frTemplate{}
	}

	return resolveMail(emailContacts, i18nTemplate, sender, recipients)
}

func resolveMail(emailContacts Contacts, i18nTemplate i18nTemplate, sender string, recipients []string) ([]byte, error) {

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

	m := subjectBody{formatRfc5322(sender), formatMultipleRfc5322(recipients), subject.String(), body.String()}

	resolvedMail := resolveTemplate(mail, m)
	return resolvedMail.Bytes(), nil
}

func resolveTemplate(template *template.Template, object interface{}) bytes.Buffer {
	var doc bytes.Buffer
	template.Execute(&doc, object)
	return doc
}

func formatRfc5322(name string) string {
	//Barry Gibbs <bg@example.com>
	return name + " <" + name + ">"
}

func formatMultipleRfc5322(names []string) string {
	rfc5322Names := ""
	for i, s := range names {
		if i > 0 {
			rfc5322Names += ", "
		}
		rfc5322Names += formatRfc5322(s)
	}
	return rfc5322Names
}
