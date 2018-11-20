package email

type enWeeklyTemplate struct {
}

func (en enWeeklyTemplate) weeklyDigestSubject() string {
	return "Anniversaire de {{.Name}}{{if yearValid .BirthDate}} -{{.Age}} an(s)- {{- end}}"
}

func (en enWeeklyTemplate) weeklyDigestBody() string {
	return "Ce sera l'anniversaire de {{.Name}} le {{formatDate .BirthDate}}. Pensez à le lui souhaiter!"
}
