package email

import "time"

type frTemplate struct {
}

func (fr frTemplate) simpleReminderSubject() string {
	return "Anniversaires du {{formatDate .RemindDate}}"
}

func (fr frTemplate) simpleReminderBody() string {
	return `Le {{formatDate .RemindDate}}, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} ans) {{- end}}
{{end}}`
}


func (fr frTemplate) weeklyDigestSubject() string {
	return "Les anniversaires de la semaine"
}

func (fr frTemplate) weeklyDigestBody() string {
	return "Ils/Elles fêteront leur anniversaire cette semaine: "
}

func (fr frTemplate) monthlyDigestSubject() string {
	return "Anniversaire de {{.Name}}{{if yearValid .BirthDate}} -{{.Age}} an(s)- {{- end}}"
}

func (fr frTemplate) monthlyDigestBody() string {
	return "Ce sera l'anniversaire de {{.Name}} le {{formatDate .BirthDate}}. Pensez à le lui souhaiter!"
}

func (fr frTemplate) dateLayout() string {
	return "02/01"
}

func (fr frTemplate) formatDate(date time.Time) string {
	return date.Format(fr.dateLayout())
}
