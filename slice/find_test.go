package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []Number
		match    matchFunc[Number]
		wantVals Number
		found    bool
	}{
		{
			name: "find 1",
			slice: []Number{
				{val: 1},
				{val: 2},
				{val: 3},
			},
			match: func(src Number) bool {
				return src.val == 1
			},
			wantVals: Number{val: 1},
			found:    true,
		},
		{
			name: "find Nothing",
			slice: []Number{
				{val: 1},
				{val: 2},
				{val: 3},
			},
			match: func(src Number) bool {
				return src.val == 4
			},
		},
		{
			name: "nil",
			match: func(src Number) bool {
				return src.val == 4
			},
			found: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, ok := Find[Number](tc.slice, tc.match)
			assert.Equal(t, tc.found, ok)
			assert.Equal(t, tc.wantVals, res)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []Number
		match    matchFunc[Number]
		wantVals []Number
		found    bool
	}{
		{
			name: "没有符合条件的",
			slice: []Number{
				{val: 1},
				{val: 3},
			},
			match: func(src Number) bool {
				return src.val%2 == 0
			},
			wantVals: []Number{},
		},
		{
			name: "找到了",
			slice: []Number{
				{val: 1},
				{val: 2},
				{val: 3},
			},
			match: func(src Number) bool {
				return src.val%2 == 1
			},
			wantVals: []Number{
				{val: 1},
				{val: 3},
			},
		},
		{
			name: "nil",
			match: func(src Number) bool {
				return src.val%2 == 1
			},
			wantVals: []Number{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FindAll[Number](tc.slice, tc.match)
			assert.Equal(t, tc.wantVals, res)
		})
	}
}

func ExampleFind() {
	val, ok := Find[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 2
	})
	fmt.Println(val, ok)
	val, ok = Find[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 4
	})
	fmt.Println(val, ok)
	// Output:
	// 2 true
	// 0 false
}

func ExampleFindAll() {
	vals := FindAll[int]([]int{2, 3, 4}, func(src int) bool {
		return src%2 == 1
	})
	fmt.Println(vals)
	vals = FindAll[int]([]int{2, 3, 4}, func(src int) bool {
		return src > 5
	})
	fmt.Println(vals)
	// Output:
	// [3]
	// []
}

type Number struct {
	val int
}
