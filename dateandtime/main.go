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

//648

func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}

	tickChannel := time.Tick(time.Second)
	index := 0

	for {
		<-tickChannel
		nameChannel <- names[index]
		index++
		if index == len(names) {
			index = 0
		}
	}
}

func main() {
	nameChannel := make(chan string)

	go writeToChannel(nameChannel)

	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
