package list

import (
	"github.com/zmsocc/generic/internal/errs"
	"github.com/zmsocc/generic/internal/slice"
)

// ArrayList 基于切片的简单封装
type ArrayList[T any] struct {
	vals []T
}

// NewArrayList 初始化一个长度为 0，容量为 cap 的 ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

// NewArrayListOf 直接使用 src，而不会执行复制
func NewArrayListOf[T any](src []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: src,
	}
}

// Get 获取索引在 index 处的值
func (a *ArrayList[T]) Get(index int) (t T, e error) {
	length := a.Len()
	if index < 0 || index >= length {
		return t, errs.NewErrIndexOutOfRange(length-1, index)
	}
	return a.vals[index], e
}

// Append 往 ArrayList 里追加数据
func (a *ArrayList[T]) Append(src ...T) error {
	a.vals = append(a.vals, src...)
	return nil
}

// Add 在 ArrayList 下标为 index 处插入元素 val
func (a *ArrayList[T]) Add(val T, index int) (err error) {
	a.vals, err = slice.Add(a.vals, val, index)
	return err
}

// Set 设置 ArrayList 下标为 index 处的值为 val
func (a *ArrayList[T]) Set(val T, index int) error {
	length := a.Len()
	if index < 0 || index >= length {
		return errs.NewErrIndexOutOfRange(length-1, index)
	}
	a.vals[index] = val
	return nil
}

// Delete 方法会在必要的时候引起缩容，其缩容规则是：
// - 如果 cap > 2048，并且长度小于等于 cap 的一半，那么就会缩容为原本的 5/8
// - 如果 64 < cap <= 2048，如果长度小于等于 cap 的 1/4，那么就会缩容为原来的 1/2
// - 如果 cap < 64，那么就不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU 去执行缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(a.vals, index)
	if err != nil {
		return t, err
	}
	a.vals = res
	a.shrink()
	return t, nil
}

// shrink 数组缩容
func (a *ArrayList[T]) shrink() {
	a.vals = slice.Shrink(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.vals))
	copy(res, a.vals)
	return res
}
