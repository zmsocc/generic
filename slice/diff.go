package slice

// DiffSet 差集，只支持 comparable 类型, src 减去 target 中与之相等的元素
// 已去重
// 并且返回值的顺序是不确定的
func DiffSet[T comparable](src []T, target []T) []T {
	srcMap := toMap(src)
	for _, v := range target {
		delete(srcMap, v)
	}
	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}
	return res
}

// DiffSetFunc 差集，支持任意类型, src 减去 target 中与之相等的元素
// 已去重
// 优先使用 DiffSet
func DiffSetFunc[T any](src []T, target []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for _, v := range src {
		if !ContainsFunc(target, func(src T) bool { return equal(src, v) }) {
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}
