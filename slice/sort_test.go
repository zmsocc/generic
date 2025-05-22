package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortAscInt(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes []int
	}{
		{
			name:    "src exist 1",
			src:     []int{1, 7, 3, 6, 5},
			wantRes: []int{1, 3, 5, 6, 7},
		},
		{
			name:    "src exist 2",
			src:     []int{-1, 7, -3, 6, 0},
			wantRes: []int{-3, -1, 0, 6, 7},
		},
		{
			name:    "src exist 3",
			src:     []int{1, 7, 3, 3, 6, 7, 3, 5},
			wantRes: []int{1, 3, 3, 3, 5, 6, 7, 7},
		},
		{
			name:    "src is empty",
			src:     []int{},
			wantRes: []int{},
		},
		{
			name:    "src has 1 element",
			src:     []int{1},
			wantRes: []int{1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SortAsc[int](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestSortAscFloat64(t *testing.T) {
	testCases := []struct {
		name    string
		src     []float64
		wantRes []float64
	}{
		{
			name:    "src exist 1",
			src:     []float64{1, 7, 3, 6, 5},
			wantRes: []float64{1, 3, 5, 6, 7},
		},
		{
			name:    "src exist 2",
			src:     []float64{-1.2, -1.5, 7, -3, 6, 0},
			wantRes: []float64{-3, -1.5, -1.2, 0, 6, 7},
		},
		{
			name:    "src exist 3",
			src:     []float64{1.234567, 7.2, 3.7, 3, 6, 7.0, 3.9, 5},
			wantRes: []float64{1.234567, 3, 3.7, 3.9, 5, 6, 7.0, 7.2},
		},
		{
			name:    "src is empty",
			src:     []float64{},
			wantRes: []float64{},
		},
		{
			name:    "src has 1 element",
			src:     []float64{1.11},
			wantRes: []float64{1.11},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SortAsc[float64](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestSortAscString(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		wantRes []string
	}{
		{
			name:    "src exist 1",
			src:     []string{"1", "a", "A"},
			wantRes: []string{"1", "A", "a"},
		},
		{
			name:    "src exist 2",
			src:     []string{"1", "3", "6.5", "6.1"},
			wantRes: []string{"1", "3", "6.1", "6.5"},
		},
		{
			name:    "src exist 3",
			src:     []string{"a", "b", "A", "Au"},
			wantRes: []string{"A", "Au", "a", "b"},
		},
		{
			name:    "src is empty",
			src:     []string{},
			wantRes: []string{},
		},
		{
			name:    "src has 1 element",
			src:     []string{"1.11"},
			wantRes: []string{"1.11"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SortAsc[string](tc.src)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
