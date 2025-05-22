package slice

// IntersectSet 取交集，只支持 comparable 类型
// 已去重
func IntersectSet[T comparable](src []T, target []T) []T {
	srcMap := toMap(src)
	res := make([]T, 0, len(target))
	for _, v := range target {
		if _, ok := srcMap[v]; ok {
			res = append(res, v)
		}
	}
	return deduplicate[T](res)
}

// IntersectSetFunc 取交集，支持任意类型
// 已去重
// 你应该优先使用 IntersectSet
func IntersectSetFunc[T any](src1 []T, src2 []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src1))
	for _, v := range src2 {
		if ContainsFunc[T](src1, func(src T) bool {
			return equal(src, v)
		}) {
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}
