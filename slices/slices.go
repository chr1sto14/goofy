package slices

import "golang.org/x/exp/maps"

func Distinct[U comparable](ss []U) []U {
	m := make(map[U]struct{})
	for _, s := range ss {
		m[s] = struct{}{}
	}
	return maps.Keys(m)
}

func Distinct2[U comparable](ss []U) []U {
	m := make(map[U]struct{})
	dd := make([]U, 0, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; ok {
			continue
		}
		m[s] = struct{}{}
		dd = append(dd, s)
	}
	return dd
}
