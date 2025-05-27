package list

import (
	"errors"
	"github.com/zmsocc/generic"
	"github.com/zmsocc/generic/internal/errs"
	"golang.org/x/exp/rand"
)

const (
	FactorP  = float32(0.25) // level i 上的结点有 FactorP 的比例出现在 level i + 1 上
	MaxLever = 32
)

// FactorP = 0.25, MaxLevel = 32 列表可包含 2^64 个元素

type skipListNode[T any] struct {
	val     T
	forward []*skipListNode[T]
}

type SkipList[T any] struct {
	head    *skipListNode[T]
	level   int // SkipList为空时, level为1
	compare generic.Comparator[T]
	length  int
}

func newSkipListNode[T any](val T, level int) *skipListNode[T] {
	return &skipListNode[T]{
		val:     val,
		forward: make([]*skipListNode[T], level),
	}
}

func NewSkipList[T any](compare generic.Comparator[T]) *SkipList[T] {
	return &SkipList[T]{
		head: &skipListNode[T]{
			forward: make([]*skipListNode[T], MaxLever),
		},
		level:   1,
		compare: compare,
	}
}

// NewSkipListOf 直接传入切片 src
func NewSkipListOf[T any](src []T, compare generic.Comparator[T]) *SkipList[T] {
	sl := NewSkipList[T](compare)
	for _, val := range src {
		sl.Insert(val)
	}
	return sl
}

// Get 获取索引为 index 处的值
func (l *SkipList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.length {
		var zero T
		return zero, errs.NewErrIndexOutOfRange(l.length-1, index)
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.forward[0]
	}
	return cur.val, nil
}

// Search 查找 SkipList 中是否含有目标值 val，有返回 true，没有返回 false
func (l *SkipList[T]) Search(val T) bool {
	cur, _ := l.traverse(val, l.level)
	cur = cur.forward[0] // 第 1 层，包含所有元素
	return cur != nil && l.compare(cur.val, val) == 0
}

// Insert 在 SkipList 中插入 val，SkipList 中的数据为增序排列(有序集合)
func (l *SkipList[T]) Insert(val T) {
	_, update := l.traverse(val, l.level)
	level := l.randomLevel()
	if level > l.level {
		for i := l.level; i < level; i++ {
			update[i] = l.head
		}
		l.level = level
	}
	newNode := newSkipListNode[T](val, level)
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	l.length++
}

func (l *SkipList[T]) DeleteElement(target T) bool {
	cur, update := l.traverse(target, l.level)
	node := cur.forward[0]
	if node == nil || l.compare(node.val, target) != 0 {
		return true
	}
	// 删除 target 节点
	for i := 0; i < l.level && update[i].forward[i] == node; i++ {
		update[i].forward[i] = node.forward[i]
	}
	// 更新层级
	for l.level > 1 && l.head.forward[l.level-1] == nil {
		l.level--
	}
	l.length--
	return true
}

func (l *SkipList[T]) Peek() (T, error) {
	cur := l.head
	cur = cur.forward[0]
	var zero T
	if cur == nil {
		return zero, errors.New("跳表为空")
	}
	return cur.val, nil
}

func (l *SkipList[T]) Len() int {
	return l.length
}

func (l *SkipList[T]) AsSlice() []T {
	cur := l.head
	slice := make([]T, 0, l.length)
	for cur.forward[0] != nil {
		slice = append(slice, cur.forward[0].val)
		cur = cur.forward[0]
	}
	return slice
}

// 查找目标值 val 的插入/删除位置，记录路径信息(update 切片)
func (l *SkipList[T]) traverse(val T, level int) (*skipListNode[T], []*skipListNode[T]) {
	update := make([]*skipListNode[T], MaxLever)
	curr := l.head
	// 从最高层向最底层逐层搜索
	for i := level - 1; i >= 0; i-- {
		// 在当前层找到最后一个小于 val 的节点
		for curr.forward[i] != nil && l.compare(curr.forward[i].val, val) < 0 {
			curr = curr.forward[i]
		}
		update[i] = curr // 记录该层的最后一个小于 val 的节点
	}
	return curr, update // 返回最终位置和路径
}

// levels的生成和跳表中元素个数无关
func (l *SkipList[T]) randomLevel() int {
	level := 1
	p := FactorP
	for (rand.Int31() & 0xFFFF) < int32(p*0xFFFF) {
		level++
	}
	if level < MaxLever {
		return level
	}
	return MaxLever
}
