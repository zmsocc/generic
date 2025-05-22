package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes []int
	}{
		{
			name:    "slice exists",
			slice:   []int{1, 2, 3, 4, 5},
			wantRes: []int{5, 4, 3, 2, 1},
		},
		{
			name:    "slice is empty",
			slice:   []int{},
			wantRes: []int{},
		},
		{
			name:    "slice is nil",
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[int](tc.slice)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestReverseSelf(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes []int
	}{
		{
			name:    "slice exists",
			slice:   []int{1, 2, 3, 4, 5},
			wantRes: []int{5, 4, 3, 2, 1},
		},
		{
			name:    "slice is empty",
			slice:   []int{},
			wantRes: []int{},
		},
		{
			name:    "slice is nil",
			slice:   nil,
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ReverseSelf[int](tc.slice)
			assert.ElementsMatch(t, tc.wantRes, tc.slice)
		})
	}
}
