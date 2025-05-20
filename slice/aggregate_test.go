package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes int
	}{
		{
			name:    "最大值",
			slice:   []int{1, -2, 3},
			wantRes: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[int](tc.slice)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []float64
		wantRes float64
	}{
		{
			name:    "最小值",
			slice:   []float64{1.1, 2.08, -0.33},
			wantRes: -0.33,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[float64](tc.slice)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []uint8
		wantRes uint8
	}{
		{
			name:    "求总和",
			slice:   []uint8{1, 2, 3},
			wantRes: 6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[uint8](tc.slice)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
