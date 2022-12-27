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
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}

	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)

	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))

	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Minutes: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Milliseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}
