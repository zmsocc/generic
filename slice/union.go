package slice

// UnionSet 并集，只支持 comparable 类型
// 已去重
// 返回值的元素顺序是不定的
func UnionSet[T comparable](src []T, target []T) []T {
	for _, v := range src {
		target = append(target, v)
	}
	return deduplicate(target)
}

// UnionSetFunc 并集，支持任意类型
// 已去重
// 优先使用 UnionSet
func UnionSetFunc[T any](src []T, target []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src)+len(target))
	res = append(res, src...)
	res = append(res, target...)
	return deduplicateFunc[T](res, equal)
}
