package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    // date="7/25/2019 13:45:00"
    layout:="1/2/2006 15:04:05"
	t,_:=time.Parse(layout, date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t:=time.Now()
    formattedTime:=t.Format("Mon, 01/02/2006, 15:04")
	return formattedTime<date
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
   t,_:=time.Parse("Monday, January 2, 2006 15:04:05", date)
    return t.Hour()>=12 && t.Hour()<18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t,_:=time.Parse("1/2/2006 15:04:05", date)
    formatted:="You have an appointment on "+t.Format("Monday, January 2, 2006, at 15:04.")
    return formatted
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date:=time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)


    return date
}
