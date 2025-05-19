package slice

import (
	"github.com/go-playground/assert/v2"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		delIndex  int
		wantSlice []int
		wantValue int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 456},
			delIndex:  0,
			wantSlice: []int{456},
			wantValue: 123,
		},
		{
			name:      "index 1",
			slice:     []int{123, 456},
			delIndex:  1,
			wantSlice: []int{123},
			wantValue: 456,
		},
		{
			name:     "index out of range",
			slice:    []int{123, 456},
			delIndex: 2,
			wantErr:  errs.NewErrIndexOutOfRange(1, 2),
		},
		{
			name:     "index smaller than 0",
			slice:    []int{123, 456},
			delIndex: -1,
			wantErr:  errs.NewErrIndexOutOfRange(1, -1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			src, res, err := Delete(tc.slice, tc.delIndex)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantValue, res)
			assert.Equal(t, tc.wantSlice, src)
		})
	}
}
