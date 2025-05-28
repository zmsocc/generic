package queue

import (
	"errors"
	"github.com/zmsocc/generic"
	"github.com/zmsocc/generic/internal/slice"
)

var (
	ErrOutOfCapacity = errors.New("generic: 超出最大容量限制")
	ErrEmptyQueue    = errors.New("generic: 队列为空")
)

type PriorityQueue[T any] struct {
	compare  generic.Comparator[T]
	capacity int
	data     []T
}

func NewPriorityQueue[T any](capacity int, compare generic.Comparator[T]) *PriorityQueue[T] {
	sliceCap := capacity
	if capacity < 1 {
		capacity = 0
		sliceCap = 64
	}
	return &PriorityQueue[T]{
		compare:  compare,
		capacity: capacity,
		data:     make([]T, 0, sliceCap),
	}
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.data)
}

func (p *PriorityQueue[T]) Cap() int {
	return p.capacity
}

func (p *PriorityQueue[T]) isEmpty() bool {
	return p.Len() < 1
}

func (p *PriorityQueue[T]) isFull() bool {
	return p.Len() == p.capacity && p.capacity > 0
}

func (p *PriorityQueue[T]) isBoundless() bool {
	return p.capacity <= 0
}

func (p *PriorityQueue[T]) Peek() (T, error) {
	if p.isEmpty() {
		var zero T
		return zero, ErrEmptyQueue
	}
	return p.data[0], nil
}

func (p *PriorityQueue[T]) Enqueue(t T) error {
	if p.isFull() {
		return ErrOutOfCapacity
	}
	p.data = append(p.data, t)
	node := len(p.data) - 1
	parent := (node - 1) / 2 // 二叉堆父子关系公式：parent = (child - 1) / 2
	for parent >= 0 && p.compare(p.data[node], p.data[parent]) < 0 {
		p.data[parent], p.data[node] = p.data[node], p.data[parent]
		node = parent
		parent = (parent - 1) >> 1
	}
	return nil
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	if p.isEmpty() {
		var zero T
		return zero, ErrEmptyQueue
	}
	pop := p.data[0]
	p.data[0] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.shrinkIfNecessary()
	p.heapSmall(p.data, len(p.data), 0)
	return pop, nil
}

func (p *PriorityQueue[T]) shrinkIfNecessary() {
	if p.isBoundless() {
		p.data = slice.Shrink[T](p.data)
	}
}

// 从节点 i 开始，向下调整以维护最小堆的性质
func (p *PriorityQueue[T]) heapSmall(data []T, n, i int) {
	minPos := i
	for {
		if left := i*2 + 1; left < n && p.compare(data[left], data[minPos]) < 0 {
			minPos = left
		}
		if right := i*2 + 2; right < n && p.compare(data[right], data[minPos]) < 0 {
			minPos = right
		}
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
		i = minPos
	}
}
