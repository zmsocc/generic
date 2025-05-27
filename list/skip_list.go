package list

import (
	"github.com/zmsocc/generic"
	"github.com/zmsocc/generic/internal/list"
)

type SkipList[T any] struct {
	skiplist *list.SkipList[T]
}

func NewSkipList[T any](compare generic.Comparator[T]) *SkipList[T] {
	pq := &SkipList[T]{}
	pq.skiplist = list.NewSkipList[T](compare)
	return pq
}

func (l *SkipList[T]) Get(index int) (T, error) {
	return l.skiplist.Get(index)
}

func (l *SkipList[T]) Search(val T) bool {
	return l.skiplist.Search(val)
}

func (l *SkipList[T]) Insert(val T) {
	l.skiplist.Insert(val)
}

func (l *SkipList[T]) DeleteElement(val T) bool {
	return l.skiplist.DeleteElement(val)
}

func (l *SkipList[T]) Len() int {
	return l.skiplist.Len()
}

func (l *SkipList[T]) Cap() int {
	return l.Len()
}

func (l *SkipList[T]) AsSlice() []T {
	return l.skiplist.AsSlice()
}
