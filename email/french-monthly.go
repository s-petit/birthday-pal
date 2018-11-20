package email

type frMonthlyTemplate struct {
}

func (fr frMonthlyTemplate) subject() string {
	return "Vos anniversaires du mois"
}

func (fr frMonthlyTemplate) body() string {
	return `Pendant le mois qui arrive, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} ans), le {{formatDate .BirthDate}} {{- end}}
{{end}}`
}
