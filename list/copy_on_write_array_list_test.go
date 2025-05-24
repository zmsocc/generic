package list

import (
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestCopyOnWriteArrayList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		src     *CopyOnWriteArrayList[int]
		index   int
		wantRes int
		wantErr error
	}{
		{
			name:    "valid index",
			src:     NewCopyOnWriteArrayListOf[int]([]int{1, 2, 3}),
			index:   2,
			wantRes: 3,
		},
		{
			name:    "invalid index",
			src:     NewCopyOnWriteArrayListOf[int]([]int{1, 2, 3}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			src:     NewCopyOnWriteArrayListOf[int]([]int{1, 2, 3}),
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.src.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestCopyOnWriteArrayList_Append(t *testing.T) {
	testCases := []struct {
		name    string
		src     *ConcurrentList[int]
		dst     []int
		wantRes []int
		wantErr error
	}{
		{
			name:    "case 1",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			dst:     []int{1, 2, 3},
			wantRes: []int{1, 2, 3, 1, 2, 3},
		},
		{
			name:    "case 2",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			dst:     nil,
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "case 3",
			src:     newConcurrentList[int](nil),
			dst:     nil,
			wantRes: []int{},
		},
		{
			name:    "case 4",
			src:     newConcurrentList[int](nil),
			dst:     []int{1, 2, 3},
			wantRes: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.src.Append(tc.dst...)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.src.AsSlice())
		})
	}
}

func TestCopyOnWriteArrayList_Add(t *testing.T) {
	testCases := []struct {
		name    string
		src     *ConcurrentList[int]
		dst     int
		index   int
		wantRes []int
		wantErr error
	}{
		{
			name:    "valid index",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			dst:     666,
			index:   1,
			wantRes: []int{1, 666, 2, 3},
		},
		{
			name:    "valid index 2",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			dst:     666,
			index:   3,
			wantRes: []int{1, 2, 3, 666},
		},
		{
			name:    "invalid index",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			dst:     666,
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			dst:     666,
			index:   4,
			wantErr: errs.NewErrIndexOutOfRange(2, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.src.Add(tc.dst, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.src.AsSlice())
		})
	}
}

func TestCopyOnWriteArrayList_Set(t *testing.T) {
	testCases := []struct {
		name    string
		src     *ConcurrentList[int]
		index   int
		dst     int
		wantRes []int
		wantErr error
	}{
		{
			name:    "set 5 by index 1",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			index:   1,
			dst:     5,
			wantRes: []int{1, 5, 3},
		},
		{
			name:    "invalid index",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			index:   -1,
			dst:     5,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			index:   4,
			dst:     5,
			wantErr: errs.NewErrIndexOutOfRange(2, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.src.Set(tc.dst, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.src.AsSlice())
		})
	}
}

func TestCopyOnWriteArrayList_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		src       *ConcurrentList[int]
		index     int
		wantRes   int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "valid index",
			src:       newConcurrentList[int]([]int{1, 2, 3}),
			index:     1,
			wantRes:   2,
			wantSlice: []int{1, 3},
		},
		{
			name:    "invalid index",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			src:     newConcurrentList[int]([]int{1, 2, 3}),
			index:   4,
			wantErr: errs.NewErrIndexOutOfRange(2, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.src.Delete(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantSlice, tc.src.AsSlice())
		})
	}
}

func TestCopyOnWriteArrayList_Len(t *testing.T) {
	testCases := []struct {
		name      string
		src       *ConcurrentList[int]
		expectLen int
	}{
		{
			name:      "与实际元素数相等",
			src:       newConcurrentList[int]([]int{1, 2, 3, 4, 5}),
			expectLen: 5,
		},
		{
			name:      "用户传入nil",
			src:       newConcurrentList[int]([]int{}),
			expectLen: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.src.Cap()
			assert.Equal(t, tc.expectLen, actual)
		})
	}
}

func TestCopyOnWriteArrayList_Cap(t *testing.T) {
	testCases := []struct {
		name      string
		src       *ConcurrentList[int]
		expectCap int
	}{
		{
			name:      "equal with actual cap",
			src:       newConcurrentList[int]([]int{1, 2, 3, 4, 5}),
			expectCap: 5,
		},
		{
			name:      "nil",
			src:       newConcurrentList[int](nil),
			expectCap: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.src.Cap()
			assert.Equal(t, tc.expectCap, res)
		})
	}
}

//func TestCopyOnWriteArrayList_AsSlice(t *testing.T) {
//	vals := []int{1, 2, 3}
//	a := newConcurrentList[int](vals)
//	slice := a.AsSlice()
//	// 内容相同
//	assert.Equal(t, vals, slice)
//	aAddr := fmt.Sprintf("%p", vals)
//	sliceAddr := fmt.Sprintf("%p", slice)
//	// 但是地址不同，也就是意味着 slice 必须是一个新创建的
//	assert.Equal(t, aAddr, sliceAddr)
//}
