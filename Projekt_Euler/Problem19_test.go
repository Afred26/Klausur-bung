package Projekt_Euler

import "fmt"

//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

func Example() {
	x := NewDate(1, 0, 2, 1904)
	fmt.Println(x)
	/*count := 0
	for x.year.year < 2001 {
		if x.day.weekday == 6 && x.day.day == 1 {
			count++
		}
		x.NextDay()
	}
	fmt.Println(count)
	*/

	// Output:
	//1
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
	case 14:
		m.numday = 29
	}

}

func (d *Date) NextDay() {
	d.day.day = d.day.day + 1
	if d.day.day > d.month.numday {
		d.day.day = 1
		d.month.month = d.month.month + 1
		if d.month.month == 13 || d.month.month == 15 {
			d.month.month = 1
			d.year.year = d.year.year + 1
		}
	}
	if d.year.leapyear {
		d.month.month = 13
	}
	d.day.weekday = (d.day.weekday + 1) % 7
}

func NewDate(d, wd, m, y int) Date {
	x := Date{day: Day{day: d, weekday: wd}, month: Month{month: m}, year: Year{year: y}}
	x.year.IsLeapYear()
	x.month.WhatNumDay()
	return x
}
