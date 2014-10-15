package main

import "fmt"

var months = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

//returns days in a month by a year
func days(mon, year int) int {
	v := months[mon]
	if year%4 == 0 && year%100 != 0 && mon == 1 {
		return v + 1
	}
	return v
}

func startSundays(day, year, n int) int {
	found := 0
	month := 0
	for i := 0; i < n; i++ { //iterate each month
		if day == 0 {
			found++
		}
		d := days(month, year)
		day = (day + d%7) % 7
		month++
		if month == 12 {
			month = 0
			year++
		}
	}
	return found
}

func main() {
	fmt.Println(startSundays(2, 1901, 12*100))
}
