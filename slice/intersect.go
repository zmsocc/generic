package slice

// IntersectSet 取交集，只支持 comparable 类型
// 已去重
func IntersectSet[T comparable](src1 []T, src2 []T) []T {
	src1Map := toMap(src1)
	res := make([]T, 0, len(src2))
	for _, v := range src2 {
		if _, ok := src1Map[v]; ok {
			res = append(res, v)
		}
	}
	return deduplicate[T](res)
}

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
