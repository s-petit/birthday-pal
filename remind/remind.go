package remind

import "time"

//Reminder holds remind context and conditions
type Reminder struct {
	CurrentDate       time.Time
	NbDaysBeforeBDay  int
	EveryDayUntilBDay bool
}

//ShouldRemind returns true when the birthdate should be reminded
func (r Reminder) ShouldRemind(birthDate time.Time) bool {
	return r.remindOnce(birthDate) || r.remindEveryDay(birthDate)
}

func (r Reminder) remindOnce(birthDate time.Time) bool {

	remindDay := r.RemindDay()

	return !r.EveryDayUntilBDay && remindDay.Day() == birthDate.Day() && remindDay.Month() == birthDate.Month()
}

func (r Reminder) remindEveryDay(birthDate time.Time) bool {

	dateAtMidnight := r.dateAtMidnight()
	remindDay := r.RemindDay()

	return r.EveryDayUntilBDay && (birthDate.Day() <= remindDay.Day() && birthDate.Month() <= remindDay.Month()) &&
		(birthDate.Day() >= dateAtMidnight.Day() && birthDate.Month() >= dateAtMidnight.Month())
}

func (r Reminder) dateAtMidnight() time.Time {
	return time.Date(r.CurrentDate.Year(), r.CurrentDate.Month(), r.CurrentDate.Day(), 0, 0, 0, 0, time.Local)
}

//RemindDay returns the day to remind
func (r Reminder) RemindDay() time.Time {
	return r.dateAtMidnight().AddDate(0, 0, r.NbDaysBeforeBDay)
}
