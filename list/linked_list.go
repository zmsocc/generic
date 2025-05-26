package list

import (
	"github.com/zmsocc/generic/internal/errs"
)

// node 双向循环链表结点
type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

// LinkedList 双向循环链表
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// NewLinkedList 创建一个双向循环链表
func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

// NewLinkedListOf 将切片转换为双向循环链表, 直接使用了切片元素的值，而没有进行复制
func NewLinkedListOf[T any](src []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	if err := list.Append(src...); err != nil {
		panic(err)
	}
	return list
}

// Get 获取索引在 index 处的链表值
func (l *LinkedList[T]) Get(index int) (T, error) {
	if !l.checkIndex(index) {
		var t T
		return t, errs.NewErrIndexOutOfRange(l.length-1, index)
	}
	n := l.findNode(index)
	return n.val, nil
}

// Append 往链表最后添加元素
func (l *LinkedList[T]) Append(src ...T) error {
	for _, t := range src {
		node := &node[T]{val: t}
		if l.length == 0 {
			node.prev = node
			node.next = node
			l.head = node
			l.tail = node
		} else {
			// 若需要高频追加操作，可缓存 head 和 tail 的指针，减少解引用次数
			tail := l.tail // 缓存 tail 指针
			head := l.head // 缓存 head 指针
			node.prev = tail
			node.next = head
			tail.next = node
			head.prev = node
			l.tail = node
		}
		l.length++
	}
	return nil
}

// Add 在 LinkedList 下标为 index 的位置插入一个元素
// 当 index 等于 LinkedList 长度等同于 Append
func (l *LinkedList[T]) Add(val T, index int) error {
	if index < 0 || index > l.length {
		return errs.NewErrIndexOutOfRange(l.length-1, index)
	}
	if l.length == index {
		return l.Append(val)
	}
	find := l.findNode(index)
	node := &node[T]{prev: find.prev, next: find, val: val}
	node.prev.next = node
	node.next.prev = node
	if index == 0 {
		l.head = node
	}
	l.length++
	return nil
}

// Set 将 LinkedList 下标为 index 的位置的元素改为 val
func (l *LinkedList[T]) Set(val T, index int) error {
	if !l.checkIndex(index) {
		return errs.NewErrIndexOutOfRange(l.length-1, index)
	}
	find := l.findNode(index)
	find.val = val
	return nil
}

// Delete 删除 LinkedList 下标在 index 处的值
func (l *LinkedList[T]) Delete(index int) error {
	if !l.checkIndex(index) {
		return errs.NewErrIndexOutOfRange(l.length-1, index)
	}
	find := l.findNode(index)
	find.prev.next = find.next
	find.next.prev = find.prev
	if index == 0 {
		l.head = find.next
	}
	l.length--
	return nil
}

// Len 返回链表长度
func (l *LinkedList[T]) Len() int {
	return l.length
}

// Cap 返回链表容量，和长度一样
func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

// AsSlice 将链表转换为切片并返回切片
func (l *LinkedList[T]) AsSlice() []T {
	if l.length == 0 {
		return nil
	}
	slice := make([]T, l.length)
	cur := l.head
	for i := 0; i < l.length; i++ {
		slice[i] = cur.val
		cur = cur.next
	}
	return slice
}

func (l *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	if index < l.length/2 {
		cur = l.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := l.length - 1; i > index; i-- {
			cur = cur.prev
		}
	}
	return cur
}

func (l *LinkedList[T]) checkIndex(index int) bool {
	return 0 <= index && index < l.Len()
}
