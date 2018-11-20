package email

type i18nTemplate interface {
	subject() string
	body() string
}
