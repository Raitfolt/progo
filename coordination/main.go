package main

import (
	"math"
	"math/rand"
	"sync"
)

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}

var squares = map[int]int{}

func calculateSquares(max, iterations int) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Cached value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value: %v = %v", val, squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}

func main() {

	numRoutines := 5
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go calculateSquares(10, 5)
	}
	waitGroup.Wait()

	Printfln("Cached values: %v", len(squares))
}
