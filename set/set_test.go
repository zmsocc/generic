package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSet_Add(t *testing.T) {
	testCases := []struct {
		name    string
		m       *MapSet[int]
		val     int
		wantRes map[int]struct{}
	}{
		{
			name:    "add val",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
			val:     3,
			wantRes: map[int]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name:    "add val 2",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
			val:     2,
			wantRes: map[int]struct{}{1: struct{}{}, 2: struct{}{}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.m.Add(tc.val)
			assert.Equal(t, tc.wantRes, tc.m.m)
		})
	}
}

func TestMapSet_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		m       *MapSet[int]
		val     int
		wantRes map[int]struct{}
	}{
		{
			name:    "delete val",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
			val:     3,
			wantRes: map[int]struct{}{1: struct{}{}, 2: struct{}{}},
		},
		{
			name:    "delete val 2",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
			val:     2,
			wantRes: map[int]struct{}{1: struct{}{}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.m.Delete(tc.val)
			assert.Equal(t, tc.wantRes, tc.m.m)
		})
	}
}

func TestMapSet_Exist(t *testing.T) {
	testCases := []struct {
		name    string
		m       *MapSet[int]
		val     int
		wantRes map[int]struct{}
		exist   bool
	}{
		{
			name:    "exist val",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
			val:     2,
			wantRes: map[int]struct{}{1: struct{}{}, 2: struct{}{}},
			exist:   true,
		},
		{
			name:    "not exist val",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
			val:     3,
			wantRes: map[int]struct{}{1: struct{}{}, 2: struct{}{}},
			exist:   false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.m.Exist(tc.val)
			assert.Equal(t, tc.exist, ok)
			assert.Equal(t, tc.wantRes, tc.m.m)
		})
	}
}

func TestMapSet_Keys(t *testing.T) {
	testCases := []struct {
		name    string
		m       *MapSet[int]
		wantRes []int
	}{
		{
			name:    "keys val",
			m:       &MapSet[int]{map[int]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			wantRes: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.m.Keys()
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
