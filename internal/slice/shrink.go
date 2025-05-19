package slice

func calCapacity(c int, l int) (int, bool) {
	if l <= 0 {
		return 0, false
	}
	if c <= 64 {
		return c, false
	} else if c <= 2048 && c/l >= 4 {
		return c / 2, true
	} else if c > 2048 && c/l >= 2 {
		factor := 0.625
		return int(float64(c) * factor), true
	}
	return c, false
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
