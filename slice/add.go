package slice

import (
	"github.com/zmsocc/generic/internal/slice"
)

// Add 在 index 处插入一个 element
func Add[T any](src []T, element T, index int) ([]T, error) {
	res, err := slice.Add[T](src, element, index)
	return res, err
}
