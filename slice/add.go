package slice

import (
	"github.com/zmsocc/generic/internal/slice"
)

func Add[T any](src []T, element T, index int) ([]T, error) {
	res, err := slice.Add[T](src, element, index)
	return res, err
}
