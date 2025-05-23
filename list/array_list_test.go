package list

import (
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestArrayList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		src     *ArrayList[int]
		index   int
		wantRes int
		wantErr error
	}{
		{
			name: "valid index",
			src: &ArrayList[int]{
				vals: []int{1, 2, 3, 4, 5},
			},
			index:   0,
			wantRes: 1,
		},
		{
			name: "invalid index",
			src: &ArrayList[int]{
				vals: []int{1, 2, 3, 4, 5},
			},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(4, -1),
		},
		{
			name:    "invalid index 2",
			src:     NewArrayListOf[int]([]int{1, 2, 3, 4, 5}),
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(4, 5),
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

func TestArrayList_Append(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		src     []int
		wantRes []int
	}{
		{
			name:    "both list and src are not empty",
			list:    NewArrayListOf[int]([]int{1, 2, 3, 4, 5}),
			src:     []int{1, 2, 3, 4, 5},
			wantRes: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
		{
			name:    "empty list, not empty src",
			list:    NewArrayListOf[int]([]int{}),
			src:     []int{1, 2, 3, 4, 5},
			wantRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "empty src, not empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			src:     []int{},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "nil src, not empty list",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			src:     nil,
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "append empty src to nil list",
			list:    NewArrayListOf[int](nil),
			src:     []int{},
			wantRes: []int{},
		},
		{
			name:    "append nil src to nil list",
			list:    NewArrayListOf[int](nil),
			src:     nil,
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Append(tc.src...)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
		})
	}
}

func TestArrayList_Add(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		val     int
		index   int
		wantRes []int
		wantErr error
	}{
		{
			name:    "valid index",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   2,
			wantRes: []int{1, 2, 666, 3},
		},
		{
			name:    "valid index 2",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   3,
			wantRes: []int{1, 2, 3, 666},
		},
		{
			name:    "invalid index",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(2, 5),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.val, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		val     int
		index   int
		wantRes []int
		wantErr error
	}{
		{
			name:    "valid index",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   2,
			wantRes: []int{1, 2, 666},
		},
		{
			name:    "valid index 2",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   0,
			wantRes: []int{666, 2, 3},
		},
		{
			name:    "invalid index",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			val:     666,
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Set(tc.val, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.list.AsSlice())
		})
	}
}

func TestArrayList_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		list      *ArrayList[int]
		index     int
		wantRes   int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "valid index",
			list:      NewArrayListOf[int]([]int{1, 2, 3}),
			index:     2,
			wantRes:   3,
			wantSlice: []int{1, 2},
		},
		{
			name:      "valid index 2",
			list:      NewArrayListOf[int]([]int{1, 2, 3}),
			index:     0,
			wantRes:   1,
			wantSlice: []int{2, 3},
		},
		{
			name:    "invalid index",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "invalid index 2",
			list:    NewArrayListOf[int]([]int{1, 2, 3}),
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.list.Delete(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantSlice, tc.list.AsSlice())
		})
	}
}

// TestArrayList_Delete_Shrinkage 测试缩容
func TestArrayList_Delete_Shrink(t *testing.T) {
	testCases := []struct {
		name    string
		cap     int // 原始容量
		length  int // 切片中的元素个数
		wantCap int // 期望缩容后的容量
	}{
		// ----- #阶段一 逻辑测试# -----
		// 只测试需求的逻辑和代码的逻辑是否一致

		// case 1: cap小于等于64，不进行缩容
		{
			name:    "case 1",
			cap:     64,
			length:  1,
			wantCap: 64,
		},
		// case 2: cap大于2048，长度小于容量的 1/2。 target:已有容量的 5/8
		{
			name:    "case 2",
			cap:     4000,
			length:  999,
			wantCap: 2500,
		},
		// case 3: cap小于等于2048，长度小于容量的四分之一。 target:缩到原本的一半
		{
			name:    "case 3",
			cap:     2048,
			length:  300,
			wantCap: 1024,
		},
		// case 4: cap > 2048，但不满足缩容条件的例子
		{
			name:    "case 4",
			cap:     4000,
			length:  3888,
			wantCap: 4000,
		},
		// case 5: cap <= 2048，但不满足缩容条件的例子
		{
			name:    "case 5",
			cap:     2048,
			length:  666,
			wantCap: 2048,
		},

		// cap <= 64，但不满足缩容条件的例子
		{
			name:    "cap <= 64",
			cap:     64,
			length:  2,
			wantCap: 64,
		},

		// ----- #阶段二 边界测试# -----
		// 测试用例边界
		// ps:测试时：
		//		（1）会默认删除一个元素，length需要+1
		//		（2）测试结果向下取整

		// case 6: cap = 65
		{
			name:    "case 6",
			cap:     65,
			length:  2,
			wantCap: 32,
		},
		{
			name:    "case 6-2",
			cap:     80,
			length:  22, // 80 * 1/4 == 20, 22 - 1 > 20, 不用缩容
			wantCap: 80,
		},
		{
			name:    "case 6-3",
			cap:     80,
			length:  21, // 80 * 1/4 == 20, 21 - 1 = 20, 要缩容
			wantCap: 40,
		},
		{
			name:    "case 7",
			cap:     2047,
			length:  10,
			wantCap: 1023, // 1023.5 ，向下取整
		},
		//  case 8 : cap 2049
		{
			name:    "case 8",
			cap:     2049,
			length:  10,
			wantCap: 1280, // 1280.625 ，向下取整
		},
		{
			name:    "case 8-1",
			cap:     2049,
			length:  1025,
			wantCap: 1280, // 二分之一为 1024，1025 删除一个元素后 1024 刚好满足
		},
		{
			name:    "case 8-2",
			cap:     2049,
			length:  1026, // 二分之一为 1024，1026 删除一个元素后 1025 刚好不满足
			wantCap: 2049,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := NewArrayList[int](tc.cap)
			for i := 0; i < tc.length; i++ {
				_ = list.Append(i)
			}
			_, _ = list.Delete(0)
			assert.Equal(t, tc.wantCap, list.Cap())
		})
	}
}
