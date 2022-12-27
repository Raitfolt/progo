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
	var d time.Duration = time.Hour + (30 * time.Minute)

	Printfln("Hours: %v", d.Hours())
	Printfln("Minutes: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Milliseconds: %v", d.Milliseconds())

	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())

	trunc := d.Truncate(time.Hour)
	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())
}
