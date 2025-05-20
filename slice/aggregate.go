package slice

import "github.com/zmsocc/generic"

// Max 求最大值
func Max[T generic.RealNumber](src []T) T {
	res := src[0]
	for i := 1; i < len(src); i++ {
		if src[i] > res {
			res = src[i]
		}
	}
	return res
}

// Min 求最小值
func Min[T generic.RealNumber](src []T) T {
	res := src[0]
	for i := 1; i < len(src); i++ {
		if src[i] < res {
			res = src[i]
		}
	}
	return res
}

// Sum 求总和
func Sum[T generic.RealNumber](src []T) T {
	sum := src[0]
	for i := 1; i < len(src); i++ {
		sum += src[i]
	}
	return sum
}
