package slices

import (
	"fmt"
	"math/rand"
	"testing"
)

func GenerateRandomString(n int) string {
	const _letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = _letters[rand.Intn(len(_letters))]
	}
	return string(out)
}

func BenchmarkDistinct(b *testing.B) {
	ttt := []struct {
		n int
	}{
		{n: 10},
		{n: 100},
		{n: 1000},
		{n: 10000},
		{n: 100000},
	}
	for _, tc := range ttt {
		b.Run(fmt.Sprintf("length %d all unique", tc.n), func(b *testing.B) {
			b.ReportAllocs()
			ss := make([]int, tc.n)
			for i := 0; i < tc.n; i++ {
				ss[i] = i
			}
			b.ResetTimer()
			Distinct(ss)
		})
		b.Run(fmt.Sprintf("length %d half same", tc.n), func(b *testing.B) {
			b.ReportAllocs()
			ss := make([]int, tc.n)
			for i := 0; i < tc.n; i++ {
				if i%2 == 0 {
					ss[i] = 2
					continue
				}
				ss[i] = i
			}
			b.ResetTimer()
			Distinct(ss)
		})
		b.Run(fmt.Sprintf("length %d all unique strings", tc.n), func(b *testing.B) {
			b.ReportAllocs()
			ss := make([]string, tc.n)
			for i := 0; i < tc.n; i++ {
				ss[i] = GenerateRandomString(32)
			}
			b.ResetTimer()
			Distinct(ss)
		})
		b.Run(fmt.Sprintf("length %d half same strings", tc.n), func(b *testing.B) {
			b.ReportAllocs()
			ss := make([]string, tc.n)
			for i := 0; i < tc.n; i++ {
				if i%2 == 0 {
					ss[i] = "same"
					continue
				}
				ss[i] = GenerateRandomString(32)
			}
			b.ResetTimer()
			Distinct(ss)
		})
	}
}
