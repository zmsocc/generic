package slice

import (
	"github.com/zmsocc/generic/internal/slice"
)

// Delete 删除索引在 index 处的元素
func Delete[T any](src []T, index int) ([]T, error) {
	src, _, err := slice.Delete[T](src, index)
	return src, err
}

// FilterDelete 删除满足条件的元素
func FilterDelete[T any](src []T, m func(src T, index int) bool) []T {
	emptyPos := 0
	for i := range src {
		// 判断是否满足删除的条件
		if m(src[i], i) {
			continue
		}
		// 移动元素
		src[emptyPos] = src[i]
		emptyPos++
	}
	return src[:emptyPos]
}
