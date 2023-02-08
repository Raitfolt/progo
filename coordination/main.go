package main

import (
	"context"
	"sync"
	"time"
)

func processRequest(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			Printfln("Stop processing - request cancelled")
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	go processRequest(ctx, &waitGroup, 10)
	time.Sleep(time.Second)
	Printfln("Canceling request")
	cancel()

	waitGroup.Wait()
}
