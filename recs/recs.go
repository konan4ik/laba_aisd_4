package recs

import (
	"strconv"
	"strings"
)

type Record struct {
	Date     Date
	FullName FullName
	Number   int
	Descrp   string
}

type Date struct {
	Year  int
	Month int
	Day   int
}

type FullName struct {
	Name      string
	SurName   string
	Otchestvo string
}

func CreateRecord(Date string, fullname string, num int, descr string) Record {
	return Record{
		Date:     CreateDate(Date),
		FullName: CreateFullName(fullname),
		Number:   num,
		Descrp:   descr,
	}
}

func CreateDate(input string) Date {
	splitted := strings.Split(input, "-")
	Year, _ := strconv.Atoi(splitted[0])
	Month, _ := strconv.Atoi(splitted[1])
	Day, _ := strconv.Atoi(splitted[2])

	return Date{
		Year: Year, Month: Month, Day: Day,
	}

}

func CreateFullName(input string) FullName {
	splitted := strings.Split(input, " ")

	return FullName{
		Name: splitted[0], SurName: splitted[1], Otchestvo: splitted[2],
	}

}
