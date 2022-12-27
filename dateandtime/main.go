package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
	//fmt.Println(label, t.Format(layout))
	//Printfln("%s: Day: %v Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
}

func main() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"

	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nyerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")
	fixZone := time.FixedZone("EDT", -4*60*60)

	if lonerr == nil && nyerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		localTime, _ := time.ParseInLocation(layout, date, local)
		fixTime, _ := time.ParseInLocation(layout, date, fixZone)

		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)
		PrintTime("Fixed:", &fixTime)
	} else {
		fmt.Println(lonerr.Error(), nyerr.Error())
	}
}
