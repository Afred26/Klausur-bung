package Projekt_Euler

import "fmt"

//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

func Example() {
	y := Year{year: 1904}
	y.IsLeapYear()
	fmt.Println(y.leapyear)
	sum := 0
	for i := 1; i <= 12; i++ {
		m := Month{month: i}
		m.WhatNumDay()
		sum = sum + m.numday
	}
	fmt.Println(sum)

	// Output:
}

type Date struct {
	day   Day
	month Month
	year  Year
}

type Day struct {
	day     int
	weekday int
}

type Month struct {
	month  int
	numday int
}

type Year struct {
	year     int
	leapyear bool
}

func (d *Year) IsLeapYear() {
	d.leapyear = false

	if d.year%4 == 0 {
		d.leapyear = true
	}
	if d.year%100 == 0 {
		d.leapyear = false
	}
	if d.year%400 == 0 {
		d.leapyear = true
	}

}

func (m *Month) WhatNumDay() {
	switch m.month {
	case 1:
		m.numday = 31
	case 2:
		m.numday = 28
	case 3:
		m.numday = 31
	case 4:
		m.numday = 30
	case 5:
		m.numday = 31
	case 6:
		m.numday = 30
	case 7:
		m.numday = 31
	case 8:
		m.numday = 31
	case 9:
		m.numday = 30
	case 10:
		m.numday = 31
	case 11:
		m.numday = 30
	case 12:
		m.numday = 31
	}

}
