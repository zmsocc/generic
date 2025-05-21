package slice

// FilterMap 执行过滤并且转化
// 如果 m 的第二个返回值是 false，那么我们会忽略第一个返回值
// 即便第二个返回值是 false，后续的元素依旧会被遍历
func FilterMap[T any, G any](src []T, m func(index int, src T) (G, bool)) []G {
	res := make([]G, 0, len(src))
	for i, v := range src {
		g, ok := m(i, v)
		if ok {
			res = append(res, g)
		}
	}
	return res
}

// 构造 map
func toMap[T comparable](src []T) map[T]struct{} {
	dataMap := make(map[T]struct{}, len(src))
	for _, v := range src {
		// 使用空结构体，减少内存消耗
		dataMap[v] = struct{}{}
	}
	return dataMap
}

// 去重
func deduplicate[T comparable](src []T) []T {
	srcMap := toMap[T](src)
	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}
	return res
}

// 去重
// 你应该优先使用 deduplicate
func deduplicateFunc[T any](src []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for i, v := range src {
		if !ContainsFunc[T](src[i+1:], func(src T) bool { return equal(src, v) }) {
			res = append(res, v)
		}
	}
	return res
}
