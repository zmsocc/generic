package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes bool
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: true,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  0,
			wantRes: false,
		},
		{
			name:    "length of slice is 0",
			slice:   []int{},
			target:  6,
			wantRes: false,
		},
		{
			name:    "slice is nil",
			target:  6,
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Contains[int](tc.slice, tc.target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  []int
		wantRes bool
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  []int{1, 2},
			wantRes: true,
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3},
			target:  []int{2, 4},
			wantRes: true,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  []int{4, 5, 6},
			wantRes: false,
		},
		{
			name:    "slice nil",
			target:  []int{1, 2, 3},
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAny[int](tc.slice, tc.target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  []int
		wantRes bool
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  []int{1, 2},
			wantRes: true,
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3},
			target:  []int{2, 4},
			wantRes: true,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  []int{4, 5, 6},
			wantRes: false,
		},
		{
			name:    "slice nil",
			target:  []int{1, 2, 3},
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAnyFunc[int](tc.slice, tc.target, func(src int, target int) bool {
				return src == target
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  []int
		wantRes bool
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  []int{1, 2},
			wantRes: true,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  []int{4, 5, 6},
			wantRes: false,
		},
		{
			name:    "target not exists 2",
			slice:   []int{1, 2, 3},
			target:  []int{2, 4},
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAll[int](tc.slice, tc.target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  []int
		wantRes bool
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  []int{1, 2},
			wantRes: true,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  []int{4, 5, 6},
			wantRes: false,
		},
		{
			name:    "target not exists 2",
			slice:   []int{1, 2, 3},
			target:  []int{2, 4},
			wantRes: false,
		},
		{
			name:    "slice nil",
			target:  []int{1, 2, 3},
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAllFunc[int](tc.slice, tc.target, func(src int, target int) bool {
				return src == target
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func ExampleContains() {
	res := Contains[int]([]int{1, 2, 3}, 3)
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsFunc() {
	res := ContainsFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 3
	})
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsAll() {
	res := ContainsAll[int]([]int{1, 2, 3}, []int{3, 1})
	fmt.Println(res)
	res = ContainsAll[int]([]int{1, 2, 3}, []int{3, 1, 4})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAllFunc() {
	res := ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1, 4}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAny() {
	res := ContainsAny[int]([]int{1, 2, 3}, []int{3, 6})
	fmt.Println(res)
	res = ContainsAny[int]([]int{1, 2, 3}, []int{4, 5, 9})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAnyFunc() {
	res := ContainsAnyFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{4, 7, 6}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}
