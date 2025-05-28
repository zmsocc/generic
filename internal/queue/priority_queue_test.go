package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zmsocc/generic"
	"testing"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		name     string
		capacity int
		data     []int
		val      int
		wantErr  error
		wantRes  []int
	}{
		{
			name:     "有界空队列",
			capacity: 10,
			data:     []int{},
			val:      10,
			wantRes:  []int{10},
		},
		{
			name:     "有界满队列",
			capacity: 6,
			data:     []int{1, 2, 3, 4, 5, 6},
			val:      10,
			wantErr:  ErrOutOfCapacity,
		},
		{
			name:     "有界非空不满队列",
			capacity: 12,
			data:     []int{6, 5, 4, 3, 2, 1},
			val:      10,
			wantRes:  []int{1, 3, 2, 6, 4, 5, 10},
		},
		{
			name:     "无界空队列",
			capacity: 0,
			data:     []int{},
			val:      10,
			wantRes:  []int{10},
		},
		{
			name:     "无界非空队列",
			capacity: 0,
			data:     []int{6, 5, 4, 3, 2, 1},
			val:      10,
			wantRes:  []int{1, 3, 2, 6, 4, 5, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := priorityQueueOf(tc.capacity, tc.data, generic.ComparatorOrdered[int])
			require.NotNil(t, p)
			err := p.Enqueue(tc.val)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.capacity, p.Cap())
			assert.Equal(t, tc.wantRes, p.data)
		})
	}
	t.Run("queue is full", func(t *testing.T) {
		p := NewPriorityQueue(3, generic.ComparatorOrdered[int])
		err := p.Enqueue(1)
		assert.NoError(t, err)
		err = p.Enqueue(2)
		assert.NoError(t, err)
		err = p.Enqueue(3)
		assert.NoError(t, err)
		err = p.Enqueue(666)
		assert.Equal(t, ErrOutOfCapacity, err)
	})
	t.Run("queue isn't full", func(t *testing.T) {
		p := NewPriorityQueue(6, generic.ComparatorOrdered[int])
		for i := p.capacity - 1; i >= 0; i-- {
			err := p.Enqueue(i)
			assert.NoError(t, err)
			fmt.Println(p.data)
		}
	})
}

func TestPriorityQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name      string
		capacity  int
		data      []int
		wantErr   error
		wantRes   int
		wantSlice []int
	}{
		{
			name:      "full queue",
			data:      []int{6, 5, 4, 3, 2, 1},
			wantRes:   1,
			wantSlice: []int{2, 3, 5, 6, 4},
		},
		{
			name:     "empty queue",
			capacity: 12,
			data:     []int{},
			wantErr:  ErrEmptyQueue,
		},
		{
			name:      "1 element in queue",
			capacity:  10,
			data:      []int{1},
			wantRes:   1,
			wantSlice: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//p := NewPriorityQueue(0, generic.ComparatorOrdered[int])
			//for _, v := range tc.data {
			//	err := p.Enqueue(v)
			//	assert.NoError(t, err)
			//}
			p := priorityQueueOf(0, tc.data, generic.ComparatorOrdered[int])
			require.NotNil(t, p)
			res, err := p.Dequeue()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantSlice, p.data)
		})
	}
	t.Run("heapSmall", func(t *testing.T) {
		p := NewPriorityQueue(6, generic.ComparatorOrdered[int])
		Res := []int{5, 3, 2, 6, 4}
		for _, v := range Res {
			err := p.Enqueue(v)
			assert.NoError(t, err)
		}
		fmt.Println(p.data)
	})
}

func TestPriorityQueue_Peek(t *testing.T) {
	testCases := []struct {
		name     string
		capacity int
		data     []int
		wantErr  error
	}{
		{
			name:     "有数据",
			capacity: 0,
			data:     []int{6, 5, 4, 3, 2, 1},
			wantErr:  ErrEmptyQueue,
		},
		{
			name:     "无数据",
			capacity: 0,
			data:     []int{},
			wantErr:  ErrEmptyQueue,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewPriorityQueue[int](tc.capacity, generic.ComparatorOrdered[int])
			for _, el := range tc.data {
				err := q.Enqueue(el)
				require.NoError(t, err)
			}
			for q.Len() > 0 {
				peek, err := q.Peek()
				assert.NoError(t, err)
				el, _ := q.Dequeue()
				assert.Equal(t, el, peek)
			}
			_, err := q.Peek()
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func priorityQueueOf(capacity int, data []int, compare generic.Comparator[int]) *PriorityQueue[int] {
	p := NewPriorityQueue(capacity, compare)
	for _, v := range data {
		err := p.Enqueue(v)
		if err != nil {
			return nil
		}
	}
	return p
}
