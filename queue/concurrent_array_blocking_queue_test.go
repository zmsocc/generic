package queue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConcurrentArrayBlockingQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		name      string
		queue     func() *ConcurrentArrayBlockingQueue[int]
		val       int
		timeout   time.Duration
		wantErr   error
		wantData  []int
		wantSlice []int
		wantLen   int
		wantHead  int
		wantTail  int
	}{
		{
			name: "empty and enqueued",
			queue: func() *ConcurrentArrayBlockingQueue[int] {
				return NewConcurrentArrayBlockingQueue[int](3)
			},
			val:       123,
			timeout:   time.Second,
			wantData:  []int{123, 0, 0},
			wantSlice: []int{123},
			wantLen:   1,
			wantTail:  1,
			wantHead:  0,
		},
		{
			name: "invalid context",
			queue: func() *ConcurrentArrayBlockingQueue[int] {
				return NewConcurrentArrayBlockingQueue[int](3)
			},
			val:       123,
			timeout:   -time.Second,
			wantData:  []int{0, 0, 0},
			wantSlice: []int{},
			wantErr:   context.DeadlineExceeded,
		},
		{
			// 入队之后就满了，恰好放在切片的最后一个位置
			name: "enqueued full last index",
			queue: func() *ConcurrentArrayBlockingQueue[int] {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				q := NewConcurrentArrayBlockingQueue[int](3)
				err := q.Enqueue(ctx, 123)
				require.NoError(t, err)
				err = q.Enqueue(ctx, 234)
				require.NoError(t, err)
				return q
			},
			val:       345,
			timeout:   time.Second,
			wantData:  []int{123, 234, 345},
			wantSlice: []int{123, 234, 345},
			wantLen:   3,
			wantTail:  0,
			wantHead:  0,
		},
		{
			// 入队之后就满了，恰好放在切片的第一个
			name: "enqueued full middle index",
			queue: func() *ConcurrentArrayBlockingQueue[int] {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				q := NewConcurrentArrayBlockingQueue[int](3)
				err := q.Enqueue(ctx, 123)
				require.NoError(t, err)
				err = q.Enqueue(ctx, 234)
				require.NoError(t, err)
				err = q.Enqueue(ctx, 345)
				require.NoError(t, err)
				val, err := q.Dequeue(ctx)
				require.NoError(t, err)
				require.Equal(t, 123, val)
				return q
			},
			val:       456,
			timeout:   time.Second,
			wantData:  []int{456, 234, 345},
			wantSlice: []int{234, 345, 456},
			wantLen:   3,
			wantTail:  1,
			wantHead:  1,
		},
		{
			// 入队之后就满了，恰好放在中间
			name: "enqueued full first index",
			queue: func() *ConcurrentArrayBlockingQueue[int] {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				q := NewConcurrentArrayBlockingQueue[int](3)
				err := q.Enqueue(ctx, 123)
				require.NoError(t, err)
				err = q.Enqueue(ctx, 234)
				require.NoError(t, err)
				err = q.Enqueue(ctx, 345)
				require.NoError(t, err)
				val, err := q.Dequeue(ctx)
				require.NoError(t, err)
				require.Equal(t, 123, val)
				val, err = q.Dequeue(ctx)
				require.NoError(t, err)
				require.Equal(t, 234, val)
				err = q.Enqueue(ctx, 456)
				require.NoError(t, err)
				return q
			},
			val:       567,
			timeout:   time.Second,
			wantData:  []int{456, 567, 345},
			wantSlice: []int{345, 456, 567},
			wantLen:   3,
			wantTail:  2,
			wantHead:  2,
		},
		{
			// 元素本身就是零值
			name: "all zero value ",
			queue: func() *ConcurrentArrayBlockingQueue[int] {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				q := NewConcurrentArrayBlockingQueue[int](3)
				err := q.Enqueue(ctx, 0)
				require.NoError(t, err)
				err = q.Enqueue(ctx, 0)
				require.NoError(t, err)
				return q
			},
			val:       0,
			timeout:   time.Second,
			wantData:  []int{0, 0, 0},
			wantSlice: []int{0, 0, 0},
			wantLen:   3,
			wantTail:  0,
			wantHead:  0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tc.timeout)
			defer cancel()
			q := tc.queue()
			err := q.Enqueue(ctx, tc.val)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantData, q.data)
			assert.Equal(t, tc.wantSlice, q.AsSlice())
			assert.Equal(t, tc.wantLen, q.Len())
			assert.Equal(t, tc.wantHead, q.head)
			assert.Equal(t, tc.wantTail, q.tail)
		})
	}
	t.Run("enqueue timeout", func(t *testing.T) {
		q := NewConcurrentArrayBlockingQueue[int](3)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err := q.Enqueue(ctx, 123)
		require.NoError(t, err)
		err = q.Enqueue(ctx, 234)
		require.NoError(t, err)
		err = q.Enqueue(ctx, 345)
		require.NoError(t, err)
		err = q.Enqueue(ctx, 456)
		require.Equal(t, context.DeadlineExceeded, err)
	})
	// 入队阻塞，而后出队，于是入队成功
	t.Run("enqueue blocking and dequeue", func(t *testing.T) {
		q := NewConcurrentArrayBlockingQueue[int](3)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		go func() {
			time.Sleep(time.Millisecond * 100)
			val, err := q.Dequeue(ctx)
			require.NoError(t, err)
			require.Equal(t, 123, val)
		}()
		err := q.Enqueue(ctx, 123)
		require.NoError(t, err)
		err = q.Enqueue(ctx, 234)
		require.NoError(t, err)
		err = q.Enqueue(ctx, 345)
		require.NoError(t, err)
		err = q.Enqueue(ctx, 456)
		require.NoError(t, err)
	})
}
