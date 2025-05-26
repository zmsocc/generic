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
