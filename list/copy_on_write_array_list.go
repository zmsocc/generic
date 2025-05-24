package list

import (
	"github.com/zmsocc/generic/internal/errs"
	"github.com/zmsocc/generic/internal/slice"
	"golang.org/x/exp/slices"
	"sync"
)

// CopyOnWriteArrayList 基于切片的简单封装，写时加锁，读不加锁，适合于读多写少的场景
type CopyOnWriteArrayList[T any] struct {
	vals  []T
	mutex *sync.Mutex
}

func NewCopyOnWriteArrayList[T any]() *CopyOnWriteArrayList[T] {
	m := &sync.Mutex{}
	return &CopyOnWriteArrayList[T]{
		vals:  make([]T, 0),
		mutex: m,
	}
}

// NewCopyOnWriteArrayListOf 直接使用 src，会执行复制
func NewCopyOnWriteArrayListOf[T any](src []T) *CopyOnWriteArrayList[T] {
	items := make([]T, len(src))
	copy(items, src)
	m := &sync.Mutex{}
	return &CopyOnWriteArrayList[T]{
		vals:  items,
		mutex: m,
	}
}

func (c *CopyOnWriteArrayList[T]) Get(index int) (t T, e error) {
	length := c.Len()
	if index < 0 || index >= length {
		return t, errs.NewErrIndexOutOfRange(length-1, index)
	}
	return c.vals[index], nil
}

func (c *CopyOnWriteArrayList[T]) Append(src []T) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newItems := slices.Clone(c.vals) // go:1.21 版本之后可用
	newItems = append(newItems, src...)
	c.vals = newItems
	return nil
}

func (c *CopyOnWriteArrayList[T]) Add(src T, index int) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newItems := slices.Clone(c.vals)
	newItems, err = slice.Add(newItems, src, index)
	if err != nil {
		return err
	}
	c.vals = newItems
	return nil
}

// Set 设置 CopyOnWriteArrayList 里 index 位置的值为 t
func (c *CopyOnWriteArrayList[T]) Set(index int, t T) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	length := c.Len()
	if index < 0 || index >= length {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	//newItems := make([]T, length)
	//copy(newItems, c.vals)
	newItems := slices.Clone(c.vals)
	newItems[index] = t
	c.vals = newItems
	return nil
}

// Delete 这里不涉及缩容，每次都是当前内容长度申请的数组容量
func (c *CopyOnWriteArrayList[T]) Delete(index int) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	length := c.Len()
	if index < 0 || index >= length {
		return errs.NewErrIndexOutOfRange(length-1, index)
	}
	newItems := slices.Clone(c.vals)
	newItems, _, err = slice.Delete(newItems, index)
	c.vals = newItems
	return err
}

func (c *CopyOnWriteArrayList[T]) Len() int {
	return len(c.vals)
}

func (c *CopyOnWriteArrayList[T]) Cap() int {
	return cap(c.vals)
}

func (c *CopyOnWriteArrayList[T]) AsSlice() []T {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	res := slices.Clone(c.vals)
	return res
}
