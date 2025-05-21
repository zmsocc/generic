package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectSet(t *testing.T) {
	testCases := []struct {
		name    string
		src1    []int
		src2    []int
		wantRes []int
	}{
		{
			name:    "intersect exists",
			src1:    []int{1, 2, 3, 4, 5},
			src2:    []int{1, 2, 3, 4},
			wantRes: []int{1, 2, 3, 4},
		},
		{
			name:    "intersect exists 2",
			src1:    []int{1, 2, 3, 4, 5},
			src2:    []int{1, 3, 0, 8, 7, 4},
			wantRes: []int{1, 3, 4},
		},
		{
			name:    "intersect exists 3",
			src1:    []int{1, 2, 3, 4, 3, 3},
			src2:    []int{1, 3, 0, 8, 3, 4},
			wantRes: []int{1, 3, 4},
		},
		{
			name:    "intersect not exists",
			src1:    []int{1, 2, 3, 4, 5},
			src2:    []int{6, 7, 8, 9, 10},
			wantRes: []int{},
		},
		{
			name:    "intersect not exists 2",
			src1:    []int{},
			src2:    []int{6, 7, 8, 9, 10},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IntersectSet(tc.src1, tc.src2)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIntersectSetFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src1    []int
		src2    []int
		wantRes []int
	}{
		{
			name:    "intersect exists",
			src1:    []int{1, 2, 3, 4, 5},
			src2:    []int{1, 2, 3, 4},
			wantRes: []int{1, 2, 3, 4},
		},
		{
			name:    "intersect exists 2",
			src1:    []int{1, 2, 3, 4, 5},
			src2:    []int{1, 3, 0, 8, 7, 4},
			wantRes: []int{1, 3, 4},
		},
		{
			name:    "intersect exists 3",
			src1:    []int{1, 2, 3, 4, 3, 3},
			src2:    []int{1, 3, 0, 8, 3, 4},
			wantRes: []int{1, 3, 4},
		},
		{
			name:    "intersect not exists",
			src1:    []int{1, 2, 3, 4, 5},
			src2:    []int{6, 7, 8, 9, 10},
			wantRes: []int{},
		},
		{
			name:    "intersect not exists 2",
			src1:    []int{},
			src2:    []int{6, 7, 8, 9, 10},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IntersectSetFunc(tc.src1, tc.src2, func(a, b int) bool { return a == b })
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
