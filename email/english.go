package email

import "time"

type enTemplate struct {
}

func (en enTemplate) subject() string {
	return "{{.Name}}'s birthday {{if yearValid .BirthDate}}-{{.Age}} yo- {{- end}}"
}

func (en enTemplate) body() string {
	return "The {{formatDate .BirthDate}} will be {{.Name}}'s birthday. Do not forget to make your wish!"
}

func (en enTemplate) dateLayout() string {
	return "01/02"
}

func (en enTemplate) formatDate(date time.Time) string {
	return date.Format(en.dateLayout())
}
