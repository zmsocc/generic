package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic/internal/errs"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name       string
		slice      []int
		addIndex   int
		addElement int
		wantSlice  []int
		wantErr    error
	}{
		{
			name:       "index 0",
			slice:      []int{123, 456},
			addIndex:   0,
			addElement: 666,
			wantSlice:  []int{666, 123, 456},
		},
		{
			name:       "index 2",
			slice:      []int{123, 456},
			addIndex:   2,
			addElement: 666,
			wantSlice:  []int{123, 456, 666},
		},
		{
			name:       "index out of range",
			slice:      []int{123, 456},
			addIndex:   -1,
			addElement: 666,
			wantErr:    errs.NewErrIndexOutOfRange(1, -1),
		},
		{
			name:       "index out of range",
			slice:      []int{123, 456},
			addIndex:   5,
			addElement: 666,
			wantErr:    errs.NewErrIndexOutOfRange(1, 5),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			src, err := Add(tc.slice, tc.addElement, tc.addIndex)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, src)
		})
	}
}

func ExampleAdd() {
	res, _ := Add[int]([]int{123, 456}, 666, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{123, 456}, 666, -1)
	fmt.Println(err)
	// output:
	// [123 456 666]
	// generic: 下标不在范围内，实际下标区间为 [0, 1], 操作的下标为 -1
}
