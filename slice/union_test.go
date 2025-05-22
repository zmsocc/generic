package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionSet(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		target  []int
		wantRes []int
	}{
		{
			name:    "normal case",
			src:     []int{1, 2, 3},
			target:  []int{1, 2, 3},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "normal case 2",
			src:     []int{1, 2, 3},
			target:  []int{4, 5, 6, 1, 3},
			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "src is empty",
			src:     []int{},
			target:  []int{1, 2, 3},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "target is empty",
			src:     []int{1, 2, 3},
			target:  []int{},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "src and target are both empty",
			src:     []int{},
			target:  []int{},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSet(tc.src, tc.target)
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}

func TestUnionSetFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		target  []int
		wantRes []int
	}{
		{
			name:    "normal case",
			src:     []int{1, 2, 3},
			target:  []int{1, 2, 3},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "normal case 2",
			src:     []int{1, 2, 3},
			target:  []int{4, 5, 6, 1, 3},
			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "src is empty",
			src:     []int{},
			target:  []int{1, 2, 3},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "target is empty",
			src:     []int{1, 2, 3},
			target:  []int{},
			wantRes: []int{1, 2, 3},
		},
		{
			name:    "src and target are both empty",
			src:     []int{},
			target:  []int{},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSetFunc[int](tc.src, tc.target, func(src int, target int) bool {
				return src == target
			})
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}
