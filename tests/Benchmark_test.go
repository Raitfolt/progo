package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	sizes := []int{10, 100, 250}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array Size %v", size), func(subB *testing.B) {
			data := make([]int, size)
			b.ResetTimer()
			for i := 0; i < subB.N; i++ {
				subB.StopTimer()
				for j := 0; j < size; j++ {
					data[j] = rand.Int()
				}
				subB.StartTimer()
				sortAndTotal(data)
			}
		})
	}
}
