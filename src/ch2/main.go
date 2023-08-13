package ch2

import (
	"fmt"
	"math/rand"
	"time"
)

func TestMain() {
	rand.Seed(time.Now().Unix())
	year := rand.Intn(1023) + 1000
	month := 2 //rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		if IsLeapYear(year) {
			daysInMonth = 29
			break
		}
		daysInMonth = 28
	case 4, 6, 7, 9, 11:
		daysInMonth = 30
	}

	fmt.Println(daysInMonth)
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(year, month, day)
}

func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}
