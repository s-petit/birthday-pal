package remind

import "time"

//Reminder holds remind context and conditions
type Reminder struct {
	//TODO SPE : La currentDate devrait etre invariable. Elle est certes pratique pour les tests, mais en vrai ca vaut toujours now.
	// a supprimer ?
	CurrentDate       time.Time
	//TODO renommer cette variable
	NbDaysBeforeBDay  int
	//TODO renommer cette variable
	EveryDayUntilBDay bool
}

//ShouldRemind returns true when the birthdate should be reminded
func (r Reminder) ShouldRemind(birthDate time.Time) bool {
	return r.remindOnce(birthDate) || r.remindEveryDay(birthDate)
}


//TODO SPE remonner
//TODO SPE plutot qu'une date, mettre un contact en parametre. Une date seule sans contexte est plus difficile a comprendre.
//TODO SPE peut etre serait-ce plus simple que ce soit le contact qui porte cette moethde, avec un reminder en parametre ?
func (r Reminder) remindOnce(birthDate time.Time) bool {

	remindDay := r.RemindDay()

	return !r.EveryDayUntilBDay && remindDay.Day() == birthDate.Day() && remindDay.Month() == birthDate.Month()
}


//TODO SPE renommer et changer le comportement
func (r Reminder) remindEveryDay(birthDate time.Time) bool {

	dateAtMidnight := r.dateAtMidnight()
	remindDay := r.RemindDay()

	return r.EveryDayUntilBDay && (birthDate.Day() <= remindDay.Day() && birthDate.Month() <= remindDay.Month()) &&
		(birthDate.Day() >= dateAtMidnight.Day() && birthDate.Month() >= dateAtMidnight.Month())
}

func (r Reminder) dateAtMidnight() time.Time {
	return time.Date(r.CurrentDate.Year(), r.CurrentDate.Month(), r.CurrentDate.Day(), 0, 0, 0, 0, time.Local)
}

//TODO SPE renommer
//RemindDay returns the day to remind
func (r Reminder) RemindDay() time.Time {
	return r.dateAtMidnight().AddDate(0, 0, r.NbDaysBeforeBDay)
}
