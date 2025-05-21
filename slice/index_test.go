package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes int
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: 2,
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3, 2, 2, 4},
			target:  2,
			wantRes: 1,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  4,
			wantRes: -1,
		},
		{
			name:    "target not exists 2",
			slice:   []int{},
			target:  4,
			wantRes: -1,
		},
		{
			name:    "slice nil",
			target:  4,
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Index[int](tc.slice, tc.target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIndexFunc(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes int
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: 2,
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3, 2, 2, 4},
			target:  2,
			wantRes: 1,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  4,
			wantRes: -1,
		},
		{
			name:    "target not exists 2",
			slice:   []int{},
			target:  4,
			wantRes: -1,
		},
		{
			name:    "slice nil",
			target:  4,
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexFunc[int](tc.slice, func(src int) bool {
				return src == tc.target
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestLastIndex(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes int
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: 2,
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3, 2, 2, 4},
			target:  2,
			wantRes: 4,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  4,
			wantRes: -1,
		},
		{
			name:    "target not exists 2",
			slice:   []int{},
			target:  4,
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := LastIndex[int](tc.slice, tc.target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestLastIndexFunc(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes int
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: 2,
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3, 2, 2, 4},
			target:  2,
			wantRes: 4,
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  4,
			wantRes: -1,
		},
		{
			name:    "target not exists 2",
			slice:   []int{},
			target:  4,
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := LastIndexFunc[int](tc.slice, func(src int) bool {
				return src == tc.target
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIndexAll(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes []int
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: []int{2},
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3, 2, 2, 4},
			target:  2,
			wantRes: []int{1, 3, 4},
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  4,
			wantRes: []int{},
		},
		{
			name:    "target not exists 2",
			slice:   []int{},
			target:  4,
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAll[int](tc.slice, tc.target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIndexAllFunc(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		target  int
		wantRes []int
	}{
		{
			name:    "target exists",
			slice:   []int{1, 2, 3},
			target:  3,
			wantRes: []int{2},
		},
		{
			name:    "target exists 2",
			slice:   []int{1, 2, 3, 2, 2, 4},
			target:  2,
			wantRes: []int{1, 3, 4},
		},
		{
			name:    "target not exists",
			slice:   []int{1, 2, 3},
			target:  4,
			wantRes: []int{},
		},
		{
			name:    "target not exists 2",
			slice:   []int{},
			target:  4,
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAllFunc[int](tc.slice, func(src int) bool {
				return src == tc.target
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
