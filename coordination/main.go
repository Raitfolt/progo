package main

import "sync"

var waitGroup = sync.WaitGroup{}

func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func main() {
	counter := 0

	waitGroup.Add(1)
	go doSum(5000, &counter)
	waitGroup.Wait()

	Printfln("Total: %v", counter)
}
