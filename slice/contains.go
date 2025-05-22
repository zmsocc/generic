package slice

// Contains 判断 src 里面是否存在 target
func Contains[T comparable](src []T, target T) bool {
	return ContainsFunc[T](src, func(src T) bool { return src == target })
}

// ContainsFunc 判断 src 里面是否存在 target
// 应该优先使用 Contains
func ContainsFunc[T any](src []T, equal func(target T) bool) bool {
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

// ContainsAny 判断 src 里面是否存在 dst 中的任何一个元素
func ContainsAny[T comparable](src []T, target []T) bool {
	srcMap := toMap[T](src)
	for _, v := range target {
		if _, ok := srcMap[v]; ok {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
// 应该优先使用 ContainsAny
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

func ContainsAll[T comparable](src []T, target []T) bool {
	srcMap := toMap[T](src)
	for _, v := range target {
		if _, ok := srcMap[v]; !ok {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断 src 里面是否存在 dst 中的所有元素
// 你应该优先使用 ContainsAll
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		if !ContainsFunc[T](src, func(src T) bool {
			return equal(src, valDst)
		}) {
			return false
		}
	}
	return true
}
