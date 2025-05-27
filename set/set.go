package set

type Set interface {
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

// Add 添加一个元素 val
func (s *MapSet[T]) Add(val T) {
	s.m[val] = struct{}{}
}

// Delete 删除值为 val 的元素
func (s *MapSet[T]) Delete(val T) {
	delete(s.m, val)
}

// Exist 判断集合中是否有 val 元素
func (s *MapSet[T]) Exist(val T) bool {
	_, ok := s.m[val]
	return ok
}

// Keys 方法返回的元素顺序不固定
func (s *MapSet[T]) Keys() []T {
	res := make([]T, 0, len(s.m))
	for k := range s.m {
		res = append(res, k)
	}
	return res
}
