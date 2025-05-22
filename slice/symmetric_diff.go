package slice

// SymmetricDiffSet 对称差集
// 已去重
// 返回值的元素顺序是不定的
func SymmetricDiffSet[T comparable](src []T, target []T) []T {
	srcMap, targetMap := toMap(src), toMap(target)
	for key := range targetMap {
		if _, ok := srcMap[key]; ok {
			delete(srcMap, key)
		} else {
			srcMap[key] = struct{}{}
		}
	}
	res := make([]T, 0, len(srcMap))
	for key := range srcMap {
		res = append(res, key)
	}
	return res
}

// SymmetricDiffSetFunc 对称差集
// 已去重
// 你应该优先使用 SymmetricDiffSet
func SymmetricDiffSetFunc[T any](src []T, target []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src)+len(target))
	for _, v := range src {
		if !ContainsFunc[T](target, func(x T) bool { return equal(v, x) }) {
			res = append(res, v)
		}
	}
	for _, v := range target {
		if !ContainsFunc[T](src, func(x T) bool { return equal(v, x) }) {
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}
