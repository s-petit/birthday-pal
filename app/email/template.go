package email

import "time"

type i18nTemplate interface {
	subject() string
	body() string
	dateLayout() string
	formatDate(date time.Time) string
}
