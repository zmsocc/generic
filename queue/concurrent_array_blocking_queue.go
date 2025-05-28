package queue

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

// ConcurrentArrayBlockingQueue 有界并发阻塞队列
type ConcurrentArrayBlockingQueue[T any] struct {
	data       []T
	mutex      *sync.RWMutex
	head       int // 队头元素下标
	tail       int // 队尾元素下标
	count      int // 包含多少元素
	enqueueCap *semaphore.Weighted
	dequeueCap *semaphore.Weighted
	zero       T // zero 不能作为返回值返回，防止用户篡改
}

// NewConcurrentArrayBlockingQueue 创建一个有界阻塞队列
// 容量会在最开始的时候就初始化好
// capacity 必须为正数
func NewConcurrentArrayBlockingQueue[T any](capacity int) *ConcurrentArrayBlockingQueue[T] {
	semaForEnqueue := semaphore.NewWeighted(int64(capacity))
	semaForDequeue := semaphore.NewWeighted(int64(capacity))
	// error暂时不处理，因为目前没办法处理，只能考虑panic掉
	// 相当于将信号量置空
	_ = semaForDequeue.Acquire(context.TODO(), int64(capacity))
	return &ConcurrentArrayBlockingQueue[T]{
		data:       make([]T, capacity),
		mutex:      &sync.RWMutex{},
		enqueueCap: semaForEnqueue,
		dequeueCap: semaForDequeue,
	}
}

// Enqueue 入队
// 通过sema来控制容量、超时、阻塞问题
func (c *ConcurrentArrayBlockingQueue[T]) Enqueue(ctx context.Context, t T) error {
	err := c.enqueueCap.Acquire(ctx, 1) // 能拿到，说明队列还有空位，可以入队，拿不到则阻塞
	if err != nil {
		return err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ctx.Err() != nil { // 拿到锁，先判断是否超时，防止在抢锁时已经超时
		c.enqueueCap.Release(1) // 超时应该主动归还信号量，避免容量泄露
		return ctx.Err()
	}
	c.data[c.tail] = t
	c.tail++
	c.count++
	if c.tail == cap(c.data) { // c.tail 已经是最后一个了，重置下标
		c.tail = 0
	}
	c.dequeueCap.Release(1) // 往出队的 sema 放入一个元素，出队的 goroutine 可以拿到并出队
	return nil
}

func (c *ConcurrentArrayBlockingQueue[T]) Dequeue(ctx context.Context) (T, error) {
	err := c.dequeueCap.Acquire(ctx, 1) // 能拿到，说明队列有元素可以取，可以出队，拿不到则阻塞
	var res T
	if err != nil {
		return res, err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ctx.Err() != nil { // 拿到锁，先判断是否超时，防止在抢锁时已经超时
		c.dequeueCap.Release(1) // 超时应该主动归还信号量，有元素消费不到
		return res, ctx.Err()
	}
	res = c.data[c.head]
	c.data[c.head] = c.zero // 为了释放内存，GC
	c.head++
	c.count--
	if c.head == cap(c.data) {
		c.head = 0
	}
	c.enqueueCap.Release(1) // 往入队的 sema 放入一个元素，入队的 goroutine 可以拿到并入队
	return res, nil
}

func (c *ConcurrentArrayBlockingQueue[T]) Len() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func (c *ConcurrentArrayBlockingQueue[T]) AsSlice() []T {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	res := make([]T, c.count)
	capacity := cap(c.data)
	for i := 0; i < c.count; i++ {
		res[i] = c.data[(i+c.head)%capacity]
	}
	return res
}
