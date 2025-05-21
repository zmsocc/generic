package slice

// Index 返回和 target 相等的第一个元素的下标
// -1 表示没找到
func Index[T comparable](src []T, target T) int {
	return IndexFunc[T](src, func(x T) bool { return x == target })
}

// IndexFunc 返回 match 返回 true 的第一个下标
// -1 表示没找到
// 你应该优先使用 Index
func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for i, v := range src {
		if match(v) {
			return i
		}
	}
	return -1
}

// LastIndex 返回和 dst 相等的最后一个元素下标
// -1 表示没找到
func LastIndex[T comparable](src []T, target T) int {
	return LastIndexFunc[T](src, func(x T) bool { return x == target })
}

// LastIndexFunc 返回和 dst 相等的最后一个元素下标
// -1 表示没找到
// 你应该优先使用 LastIndex
func LastIndexFunc[T any](src []T, match matchFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

// IndexAll 返回和 dst 相等的所有元素的下标
func IndexAll[T comparable](src []T, target T) []int {
	return IndexAllFunc(src, func(x T) bool { return x == target })
}

// IndexAllFunc 返回和 match 返回 true 的所有元素的下标
// 你应该优先使用 IndexAll
func IndexAllFunc[T any](src []T, match matchFunc[T]) []int {
	res := make([]int, 0, len(src))
	for i, v := range src {
		if match(v) {
			res = append(res, i)
		}
	}
	return res
}
