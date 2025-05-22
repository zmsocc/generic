package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymmetricDiffSet(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		target  []int
		wantRes []int
	}{
		{
			name:    "no intersect",
			src:     []int{1, 2, 3},
			target:  []int{4, 5, 6},
			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "src contains target",
			src:     []int{1, 2, 3},
			target:  []int{1, 2},
			wantRes: []int{3},
		},
		{
			name:    "target contains src",
			src:     []int{4, 5},
			target:  []int{4, 5, 6},
			wantRes: []int{6},
		},
		{
			name:    "equal",
			src:     []int{1, 2, 3},
			target:  []int{1, 2, 3},
			wantRes: []int{},
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
			name:    "both src and target are empty",
			src:     []int{},
			target:  []int{},
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			target:  []int{4, 5, 6},
			wantRes: []int{4, 5, 6},
		},
		{
			name:    "dst nil",
			src:     []int{4, 5, 6},
			target:  nil,
			wantRes: []int{4, 5, 6},
		},
		{
			name:    "both nil",
			src:     nil,
			target:  nil,
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SymmetricDiffSet(tc.src, tc.target)
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}

func TestSymmetricDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		target  []int
		wantRes []int
	}{
		{
			name:    "no intersect",
			src:     []int{1, 3, 2},
			target:  []int{4, 5, 6},
			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "src contains target",
			src:     []int{1, 2, 3},
			target:  []int{1, 2},
			wantRes: []int{3},
		},
		{
			name:    "target contains src",
			src:     []int{4, 5},
			target:  []int{4, 5, 6},
			wantRes: []int{6},
		},
		{
			name:    "equal",
			src:     []int{1, 2, 3},
			target:  []int{1, 2, 3},
			wantRes: []int{},
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
			name:    "both src and target are empty",
			src:     []int{},
			target:  []int{},
			wantRes: []int{},
		},
		{
			name:    "src nil",
			src:     nil,
			target:  []int{4, 5, 6},
			wantRes: []int{4, 5, 6},
		},
		{
			name:    "dst nil",
			src:     []int{4, 5, 6},
			target:  nil,
			wantRes: []int{4, 5, 6},
		},
		{
			name:    "both nil",
			src:     nil,
			target:  nil,
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SymmetricDiffSetFunc(tc.src, tc.target, func(src int, target int) bool {
				return src == target
			})
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}
