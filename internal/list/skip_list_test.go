package list

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestSkipList_Get(t *testing.T) {
	testCases := []struct {
		name      string
		compare   generic.Comparator[int]
		list      *SkipList[int]
		index     int
		wantRes   int
		wantSlice []int
		wantSize  int
		wantErr   error
	}{
		{
			name:      "get index 1 from [1, 3]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3}, generic.ComparatorOrdered[int]),
			index:     1,
			wantRes:   3,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantErr:   nil,
		},
		{
			name:      "get index 2 from [1, 3, 5, 6, 7, 2, 9]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			index:     2,
			wantRes:   3,
			wantSlice: []int{1, 2, 3, 5, 6, 7, 9},
			wantSize:  7,
			wantErr:   nil,
		},
		{
			name:    "get index -1 from [1, 3, 5, 6, 7, 2, 9]",
			compare: generic.ComparatorOrdered[int],
			list:    NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(6, -1),
		},
		{
			name:    "get index 7 from [1, 3, 5, 6, 7, 2, 9]",
			compare: generic.ComparatorOrdered[int],
			list:    NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			index:   7,
			wantErr: errs.NewErrIndexOutOfRange(6, 7),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.list.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantSize, tc.list.length)
			assert.Equal(t, tc.wantSlice, tc.list.AsSlice())
		})
	}
}

func TestSkipList_Search(t *testing.T) {
	testCases := []struct {
		name      string
		list      *SkipList[int]
		compare   generic.Comparator[int]
		val       int
		wantRes   bool
		wantSize  int
		wantSlice []int
	}{
		{
			name:      "search 2 from [1, 2, 3]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 2, 3}, generic.ComparatorOrdered[int]),
			val:       2,
			wantRes:   true,
			wantSize:  3,
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "search 2 from [1,3]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3}, generic.ComparatorOrdered[int]),
			val:       2,
			wantRes:   false,
			wantSize:  2,
			wantSlice: []int{1, 3},
		},
		{
			name:      "search 2 from []",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{}, generic.ComparatorOrdered[int]),
			val:       2,
			wantRes:   false,
			wantSize:  0,
			wantSlice: []int{},
		},
		{
			name:      "search 2 from [1, 3, 5, 6, 7, 2, 9]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       2,
			wantRes:   true,
			wantSize:  7,
			wantSlice: []int{1, 2, 3, 5, 6, 7, 9},
		},
		{
			name:      "search 2 from [5, 3, 5, 6, 7, 2, 9]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{5, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       5,
			wantRes:   true,
			wantSize:  7,
			wantSlice: []int{2, 3, 5, 5, 6, 7, 9},
		},
		{
			name:      "search 2 from [1, 3, 5, 6, 7, 2, 9]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       10,
			wantRes:   false,
			wantSize:  7,
			wantSlice: []int{1, 2, 3, 5, 6, 7, 9},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.list.Search(tc.val)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantSize, tc.list.length)
			assert.Equal(t, tc.wantSlice, tc.list.AsSlice())
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	testCases := []struct {
		name      string
		list      *SkipList[int]
		compare   generic.Comparator[int]
		val       int
		wantSize  int
		wantSlice []int
	}{
		{
			name:      "insert 2 into [1, 3]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3}, generic.ComparatorOrdered[int]),
			val:       2,
			wantSlice: []int{1, 2, 3},
			wantSize:  3,
		},
		{
			name:      "insert 2 into [1, 2, 3]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 2, 3}, generic.ComparatorOrdered[int]),
			val:       2,
			wantSlice: []int{1, 2, 2, 3},
			wantSize:  4,
		},
		{
			name:      "insert 200 into [1, 3, 5, 6, 7, 2, 9]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       200,
			wantSlice: []int{1, 200},
			wantSize:  8,
		},
		{
			name:      "insert -1 into [1, 3, 5, 6, 7, 2, 9]",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       -1,
			wantSlice: []int{-1, 1, 2, 3, 5, 6, 7, 9},
			wantSize:  8,
		},
		{
			name:      "insert 200 into []",
			compare:   generic.ComparatorOrdered[int],
			list:      NewSkipListOf[int]([]int{}, generic.ComparatorOrdered[int]),
			val:       200,
			wantSlice: []int{200},
			wantSize:  1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Insert(tc.val)
			assert.Equal(t, tc.wantSlice, tc.list.AsSlice())
			assert.Equal(t, tc.wantSize, tc.list.length)
		})
	}
}

func TestSkipList_DeleteElement(t *testing.T) {
	testCases := []struct {
		name      string
		list      *SkipList[int]
		val       int
		wantSize  int
		wantRes   bool
		wantSlice []int
	}{
		{
			name:      "delete 2 from [1, 3, 5, 6, 7, 2, 9]",
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       2,
			wantRes:   true,
			wantSize:  6,
			wantSlice: []int{1, 3, 5, 6, 7, 9},
		},
		{
			name:      "delete 200 from [1, 3, 5, 6, 7, 2, 9]",
			list:      NewSkipListOf([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       200,
			wantRes:   true,
			wantSize:  7,
			wantSlice: []int{1, 2, 3, 5, 6, 7, 9},
		},
		{
			name:      "delete 1 from []",
			list:      NewSkipListOf[int]([]int{}, generic.ComparatorOrdered[int]),
			val:       1,
			wantRes:   true,
			wantSize:  0,
			wantSlice: []int{},
		},
		{
			name:      "delete 7 from [1, 3, 5, 6, 7, 2, 9]",
			list:      NewSkipListOf[int]([]int{1, 3, 5, 6, 7, 2, 9}, generic.ComparatorOrdered[int]),
			val:       7,
			wantRes:   true,
			wantSize:  6,
			wantSlice: []int{1, 2, 3, 5, 6, 9},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.list.DeleteElement(tc.val)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantSize, tc.list.length)
			assert.Equal(t, tc.wantSlice, tc.list.AsSlice())
		})
	}
}

func TestSkipList_Peek(t *testing.T) {
	testCases := []struct {
		name      string
		list      *SkipList[int]
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "peek [1,3]",
			list:      NewSkipListOf[int]([]int{1, 3}, generic.ComparatorOrdered[int]),
			wantSlice: []int{1, 3},
			wantVal:   1,
			wantErr:   nil,
		},
		{
			name:      "peek []",
			list:      NewSkipListOf[int]([]int{}, generic.ComparatorOrdered[int]),
			wantSlice: []int{},
			wantVal:   0,
			wantErr:   errors.New("跳表为空"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.list.Peek()
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}
