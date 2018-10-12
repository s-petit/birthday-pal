package email

import "time"

type i18nTemplate interface {
	simpleReminderSubject() string
	simpleReminderBody() string
	weeklyDigestSubject() string
	weeklyDigestBody() string
	monthlyDigestSubject() string
	monthlyDigestBody() string
	dateLayout() string
	formatDate(date time.Time) string
}
