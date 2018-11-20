package email

type enMonthly struct {
}

func (en enMonthly) subject() string {
	return "Anniversaire de {{.Name}}{{if yearValid .BirthDate}} -{{.Age}} an(s)- {{- end}}"
}

func (en enMonthly) body() string {
	return "Ce sera l'anniversaire de {{.Name}} le {{formatDate .BirthDate}}. Pensez à le lui souhaiter!"
}
