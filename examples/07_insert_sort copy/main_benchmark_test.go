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
	b.Run("insert binary", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(args, data[:])
			b.StartTimer()
			InsertSort(args, BinSearch)
		}
	})
	b.Run("insert exponential", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(args, data[:])
			b.StartTimer()
			InsertSort(args, ExpSearch)
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
