package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		delIndex  int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 456},
			delIndex:  0,
			wantSlice: []int{456},
		},
		{
			name:     "index out of range",
			slice:    []int{123, 456},
			delIndex: 2,
			wantErr:  errs.NewErrIndexOutOfRange(1, 2),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete(tc.slice, tc.delIndex)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleDelete() {
	src, _ := Delete[int]([]int{123, 456}, 0)
	fmt.Println(src)
	_, err := Delete[int]([]int{123, 456}, 2)
	fmt.Println(err)
	// output:
	// [456]
	// generic: 下标不在范围内，实际下标区间为 [0, 1], 操作的下标为 2
}

func TestFilterDelete(t *testing.T) {
	testCases := []struct {
		name         string
		slice        []int
		delCondition func(slice int, index int) bool
		wantSlice    []int
	}{
		{
			name:  "空切片",
			slice: []int{},
			delCondition: func(slice int, index int) bool {
				return false
			},
			wantSlice: []int{},
		},
		{
			name:  "不删除元素",
			slice: []int{123, 456},
			delCondition: func(slice int, index int) bool {
				return false
			},
			wantSlice: []int{123, 456},
		},
		{
			name:  "删除下标为偶数的元素",
			slice: []int{123, 456, 789},
			delCondition: func(slice int, index int) bool {
				return index%2 == 0
			},
			wantSlice: []int{456},
		},
		{
			name:  "删除下标为奇数的元素",
			slice: []int{123, 456, 789},
			delCondition: func(slice int, index int) bool {
				return index%2 == 1
			},
			wantSlice: []int{123, 789},
		},
		{
			name:  "删除奇数的元素",
			slice: []int{123, 456, 789},
			delCondition: func(slice int, index int) bool {
				return slice%2 == 1
			},
			wantSlice: []int{456},
		},
		{
			name:  "删除所有元素",
			slice: []int{123, 456, 789},
			delCondition: func(slice int, index int) bool {
				return true
			},
			wantSlice: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterDelete(tc.slice, tc.delCondition)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleFilterDelete() {
	res := FilterDelete[int]([]int{123, 456, 789}, func(src int, i int) bool {
		return i%2 == 0
	})
	fmt.Println(res)
	// output:
	// [456]
}
