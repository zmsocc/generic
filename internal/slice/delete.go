package slice

import "github.com/zmsocc/generic/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length-1, index)
	}
	res := src[index]
	src = append(src[:index], src[index+1:]...)
	return src, res, nil
}
