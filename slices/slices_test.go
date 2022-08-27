package slices

import (
	"testing"
)

func BenchmarkDistinct(b *testing.B) {
	ttt := []struct {
		n    int
		name string
	}{
		{n: 10, name: "length 10"},
		{n: 100, name: "length 100"},
		{n: 1000, name: "length 1000"},
		{n: 10000, name: "length 10000"},
		{n: 100000, name: "length 100000"},
	}
	for _, tc := range ttt {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			ss := make([]int, tc.n)
			for i := 0; i < tc.n; i++ {
				ss[i] = i
			}
			b.ResetTimer()
			Distinct(ss)
		})
	}
}

func BenchmarkDistinct2(b *testing.B) {
	ttt := []struct {
		n    int
		name string
	}{
		{n: 10, name: "length 10"},
		{n: 100, name: "length 100"},
		{n: 1000, name: "length 1000"},
		{n: 10000, name: "length 10000"},
		{n: 100000, name: "length 100000"},
	}
	for _, tc := range ttt {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			ss := make([]int, tc.n)
			for i := 0; i < tc.n; i++ {
				ss[i] = i
			}
			b.ResetTimer()
			Distinct2(ss)
		})
	}
}
