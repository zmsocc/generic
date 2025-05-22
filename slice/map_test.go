package slice

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFilterMap(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes []string
	}{
		{
			name:    "slice exists",
			slice:   []int{1, 2, 3},
			wantRes: []string{"1", "2", "3"},
		},
		{
			name:    "slice has element",
			slice:   []int{1, -2, 3},
			wantRes: []string{"1", "3"},
		},
		{
			name:    "slice is empty",
			slice:   []int{},
			wantRes: []string{},
		},
		{
			name:    "slice is nil",
			wantRes: []string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterMap[int, string](tc.slice, func(index int, src int) (string, bool) {
				return strconv.Itoa(src), src >= 0
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
