package email

import "time"

type i18nDateTemplate interface {
	dateLayout() string
	formatDate(date time.Time) string
}
