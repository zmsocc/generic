package list

import (
	"sync"
)

type ConcurrentList[T any] struct {
	List[T]
	lock sync.RWMutex
}

func (c *ConcurrentList[T]) Get(index int) (T, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Get(index)
}

func (c *ConcurrentList[T]) Append(src ...T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Append(src...)
}

func (c *ConcurrentList[T]) Add(val T, index int) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Add(val, index)
}

func (c *ConcurrentList[T]) Set(val T, index int) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Set(val, index)
}

func (c *ConcurrentList[T]) Delete(index int) (T, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.List.Delete(index)
}

func (c *ConcurrentList[T]) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Len()
}

func (c *ConcurrentList[T]) Cap() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.Cap()
}

func (c *ConcurrentList[T]) AsSlice() []T {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.List.AsSlice()
}
