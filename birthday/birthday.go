package birthday

import "time"


func Remind(now time.Time, birthday time.Time, nbDaysBefore int) bool {

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	delay := midnight.AddDate(0, 0, nbDaysBefore)

	return delay.Day() == birthday.Day() && delay.Month() == birthday.Month()
}
