package slice

import "github.com/zmsocc/generic/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length-1, index)
	}

	// 先将 src 扩展为一个元素
	var zeroVal T
	src = append(src, zeroVal)
	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}
