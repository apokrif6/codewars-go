//You have an amount of money a0 > 0 and you deposit it with an interest rate of p percent divided by 360 per day on the 1st of January 2016. You want to have an amount a >= a0.
Task:

//The function date_nb_days (or dateNbDays...) with parameters a0, a, p will return, as a string, the date on which you have just reached a.
//Example:

//date_nb_days(100, 101, 0.98) --> "2017-01-01" (366 days)

//date_nb_days(100, 150, 2.00) --> "2035-12-26" (7299 days)
//Notes:

//    The return format of the date is "YYYY-MM-DD"
//    The deposit is always on the "2016-01-01"
//    Don't forget to take the rate for a day to be p divided by 36000 since banks consider that there are 360 days in a year.
//    You have: a0 > 0, p% > 0, a >= a0

package kata

import "time"

func DateNbDays(a0 int, a int, p int) string {
	have := float64(a0)
	days := 0
	rate := float64(p) / 36000
  
	for ; have < float64(a); days++ {
		have += have * rate
	}
  
	data := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	targetDate := data.Add(time.Duration(days) * time.Hour * 24)
  
	return targetDate.Format("2006-01-02")
}
