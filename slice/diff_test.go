package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffSet(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		target  []int
		wantRes []int
	}{
		{
			name:    "normal case",
			src:     []int{1, 2, 3, 4, 5},
			target:  []int{1, 3, 7, 2},
			wantRes: []int{4, 5},
		},
		{
			name:    "src less than target",
			src:     []int{1, 2, 3},
			target:  []int{1, 2, 3, 4},
			wantRes: []int{},
		},
		{
			name:    "src has duplicate elements",
			src:     []int{1, 2, 3, 4, 5, 4, 4},
			target:  []int{1, 2, 3},
			wantRes: []int{4, 5},
		},
		{
			name:    "src equal target",
			src:     []int{1, 2, 3, 4, 5},
			target:  []int{1, 3, 4, 2, 5},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSet(tc.src, tc.target)
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}

func TestDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		target  []int
		wantRes []int
	}{
		{
			name:    "normal case",
			src:     []int{1, 2, 3, 4, 5},
			target:  []int{1, 3, 7, 2},
			wantRes: []int{4, 5},
		},
		{
			name:    "src less than target",
			src:     []int{1, 2, 3},
			target:  []int{1, 2, 3, 4},
			wantRes: []int{},
		},
		{
			name:    "src has duplicate elements",
			src:     []int{1, 2, 3, 4, 5, 4, 4},
			target:  []int{1, 2, 3},
			wantRes: []int{4, 5},
		},
		{
			name:    "src equal target",
			src:     []int{1, 2, 3, 4, 5},
			target:  []int{1, 3, 4, 2, 5},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSetFunc[int](tc.src, tc.target, func(src int, target int) bool {
				return src == target
			})
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}
