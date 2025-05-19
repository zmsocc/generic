package slice

import (
	"github.com/go-playground/assert/v2"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addValue  int
		addIndex  int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 456},
			addValue:  666,
			addIndex:  0,
			wantSlice: []int{666, 123, 456},
		},
		{
			name:      "index middle",
			slice:     []int{123, 456},
			addValue:  666,
			addIndex:  1,
			wantSlice: []int{123, 666, 456},
		},
		{
			name:      "index last",
			slice:     []int{123, 456},
			addValue:  666,
			addIndex:  2,
			wantSlice: []int{123, 456, 666},
		},
		{
			name:     "index out of range",
			slice:    []int{123, 456},
			addValue: 666,
			addIndex: 3,
			wantErr:  errs.NewErrIndexOutOfRange(1, 3),
		},
		{
			name:     "index smaller than 0",
			slice:    []int{123, 456},
			addValue: 666,
			addIndex: -1,
			wantErr:  errs.NewErrIndexOutOfRange(1, -1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addValue, tc.addIndex)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
