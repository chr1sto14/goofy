package slices

// Distinct returns the unique values in ss.
// order not preserved.
func Distinct[U comparable](ss []U) []U {
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
