package slice

import "github.com/zmsocc/generic"

type Ordered interface {
	generic.Ordered
}

// SortAsc 使用快速排序算法，实现升序排列
func SortAsc[T Ordered](src []T) []T {
	// 创建副本避免修改原切片
	sorted := make([]T, len(src))
	copy(sorted, src)
	// 实现快速排序算法
	quickSort[T](sorted, 0, len(sorted)-1)
	return sorted
}

// 快速排序实现
func quickSort[T Ordered](src []T, low, high int) {
	if low < high {
		pi := partition[T](src, low, high)
		quickSort[T](src, low, pi-1)
		quickSort[T](src, pi+1, high)
	}
}

// 分区函数
func partition[T Ordered](src []T, low, high int) int {
	pivot := src[high]
	i := low - 1
	for j := low; j < high; j++ {
		if src[j] < pivot {
			i++
			src[i], src[j] = src[j], src[i]
		}
	}
	src[i+1], src[high] = src[high], src[i+1]
	return i + 1
}
