package list

import (
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		src     []int
		wantRes *LinkedList[int]
		wantErr error
	}{
		{
			name:    "append non-empty values to non-empty list",
			list:    NewLinkedListOf[int]([]int{123}),
			src:     []int{234, 456},
			wantRes: NewLinkedListOf[int]([]int{123, 234, 456}),
		},
		{
			name:    "append empty values to non-empty list",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			src:     []int{},
			wantRes: NewLinkedListOf[int]([]int{1, 2, 3}),
		},
		{
			name:    "append nil to non-empty list",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			src:     nil,
			wantRes: NewLinkedListOf[int]([]int{1, 2, 3}),
		},
		{
			name:    "append non-empty values to empty list",
			list:    NewLinkedListOf[int]([]int{}),
			src:     []int{1, 2, 3},
			wantRes: NewLinkedListOf[int]([]int{1, 2, 3}),
		},
		{
			name:    "append non-empty values to nil list",
			list:    NewLinkedListOf[int](nil),
			src:     []int{1, 2, 3},
			wantRes: NewLinkedListOf[int]([]int{1, 2, 3}),
		},
		{
			name:    "append empty values to nil list",
			list:    NewLinkedListOf[int](nil),
			src:     []int{},
			wantRes: NewLinkedListOf[int]([]int{}),
		},
		{
			name:    "append nil to nil list",
			list:    NewLinkedListOf[int](nil),
			src:     nil,
			wantRes: NewLinkedListOf[int](nil),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Append(tc.src...)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantRes.AsSlice(), tc.list.AsSlice())
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		index   int
		wantRes int
		wantErr error
	}{
		{
			name:    "get non-existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "get non-existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
		{
			name:    "get existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   0,
			wantRes: 1,
		},
		{
			name:    "get existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   2,
			wantRes: 3,
		},
		{
			name:    "get existent index 3",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   1,
			wantRes: 2,
		},
		{
			name:    "empty list",
			list:    NewLinkedListOf[int](nil),
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(-1, 0),
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
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		index   int
		value   int
		wantRes *LinkedList[int]
		wantErr error
	}{
		{
			name:    "add non-existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   -1,
			value:   666,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "add existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   2,
			value:   666,
			wantRes: NewLinkedListOf[int]([]int{1, 2, 666, 3}),
		},
		{
			name:    "add existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   3,
			value:   666,
			wantRes: NewLinkedListOf[int]([]int{1, 2, 3, 666}),
		},
		{
			name:    "add existent index 3",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   0,
			value:   666,
			wantRes: NewLinkedListOf[int]([]int{666, 1, 2, 3}),
		},
		{
			name:    "add non-existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   4,
			wantErr: errs.NewErrIndexOutOfRange(2, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.value, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes.AsSlice(), tc.list.AsSlice())
		})
	}
}

func TestLinkedList_Set(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		index   int
		value   int
		wantRes *LinkedList[int]
		wantErr error
	}{
		{
			name:    "set non-existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   -1,
			value:   666,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "add existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   2,
			value:   666,
			wantRes: NewLinkedListOf[int]([]int{1, 2, 666}),
		},
		{
			name:    "add existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   3,
			value:   666,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
		{
			name:    "add existent index 3",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   1,
			value:   666,
			wantRes: NewLinkedListOf[int]([]int{1, 666, 3}),
		},
		{
			name:    "add non-existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   4,
			wantErr: errs.NewErrIndexOutOfRange(2, 4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Set(tc.value, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes.AsSlice(), tc.list.AsSlice())
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		list    *LinkedList[int]
		index   int
		wantRes *LinkedList[int]
		wantErr error
	}{
		{
			name:    "delete non-existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "delete non-existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
		{
			name:    "delete existent index",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   0,
			wantRes: NewLinkedListOf[int]([]int{2, 3}),
		},
		{
			name:    "delete existent index 2",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   1,
			wantRes: NewLinkedListOf[int]([]int{1, 3}),
		},
		{
			name:    "delete existent index 3",
			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
			index:   2,
			wantRes: NewLinkedListOf[int]([]int{1, 2}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Delete(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes.AsSlice(), tc.list.AsSlice())
		})
	}
}
