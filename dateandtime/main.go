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

func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}

	ticker := time.NewTicker(time.Second)
	index := 0

	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
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
