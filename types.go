package generic

// Comparator
// 0 表示 src = dst
// -1 表示 src < dst
// 1 表示 src > dst
type Comparator[T any] func(src T, dst T) int

func ComparatorOrdered[T Ordered](src T, dst T) int {
	if src == dst {
		return 0
	} else if src < dst {
		return -1
	}
	return 1
}
