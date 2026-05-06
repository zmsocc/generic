package slice

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestShrink(t *testing.T) {
	testCases := []struct {
		name        string
		originCap   int
		enqueueLoop int
		wantCap     int
	}{
		{
			name:        "less than 64",
			originCap:   32,
			enqueueLoop: 6,
			wantCap:     32,
		},
		{
			name:        "between 64 and 2048, AND c/l < 4 ",
			originCap:   666,
			enqueueLoop: 333,
			wantCap:     666,
		},
		{
			name:        "between 64 and 2048, AND c/l >= 4 ",
			originCap:   666,
			enqueueLoop: 111,
			wantCap:     333,
		},
		{
			name:        "more than 2048, AND c/l < 2 ",
			originCap:   3000,
			enqueueLoop: 2000,
			wantCap:     3000,
		},
		{
			name:        "more than 2048, AND c/l >= 2 ",
			originCap:   3000,
			enqueueLoop: 1000,
			wantCap:     3000 * 0.625,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]int, 0, tc.originCap)
			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink[int](l)
			assert.Equal(t, tc.wantCap, cap(l))
		})
	}
}
