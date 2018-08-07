package hexe

// stepValue transforms an index to a value in [min, max].
//
// Let there be the set V, which contains n values from [min, max],
// where the difference between all values is equal. The returned value is
// the element of the set V with the index of i.
//
// Requirements: n > 0 && i < n && max >= min
// If the requirements are violated, this function panics.
//
func stepValue(i, n uint, min, max float64) float64 {
	if n == 0 || i == n || min > max {
		panic("domain requirements are violated")
	}
	step := ((max - min) / float64(n-1))
	return float64(i)*step + min
}
