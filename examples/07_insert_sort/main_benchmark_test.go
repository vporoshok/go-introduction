package main

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	const size = 10000
	data := [size]int{}
	for i := range data {
		data[i] = rand.Int()
	}
	args := make([]int, len(data))
	b.Run("insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(args, data[:])
			b.StartTimer()
			InsertSort(args)
		}
	})
	b.Run("standard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(args, data[:])
			b.StartTimer()
			sort.Ints(args)
		}
	})
}
