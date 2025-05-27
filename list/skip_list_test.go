package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zmsocc/generic"
	"github.com/zmsocc/generic/internal/list"
	"testing"
)

func TestSkipList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		compare generic.Comparator[int]
		index   int
		wantRes int
	}{
		{
			name:    "no err is ok",
			compare: generic.ComparatorOrdered[int],
			index:   0,
			wantRes: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			res, err := sl.Get(tc.index)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestSkipList_Search(t *testing.T) {
	testCases := []struct {
		name    string
		compare generic.Comparator[int]
		val     int
		wantRes bool
	}{
		{
			name:    "no err is ok",
			compare: generic.ComparatorOrdered[int],
			val:     1,
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			ok := sl.Search(tc.val)
			assert.Equal(t, tc.wantRes, ok)
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	testCases := []struct {
		name      string
		compare   generic.Comparator[int]
		key       int
		wantSlice []int
	}{
		{
			name:      "no err is ok",
			compare:   generic.ComparatorOrdered[int],
			key:       -1,
			wantSlice: []int{-1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			sl.Insert(tc.key)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
		})
	}
}

func TestSkipList_DeleteElement(t *testing.T) {
	testCases := []struct {
		name     string
		compare  generic.Comparator[int]
		value    int
		wantBool bool
	}{
		{
			name:     "no err is ok",
			compare:  generic.ComparatorOrdered[int],
			value:    -1,
			wantBool: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			ok := sl.DeleteElement(tc.value)
			assert.Equal(t, tc.wantBool, ok)
		})
	}
}

func TestSkipList_Len(t *testing.T) {
	testCases := []struct {
		name     string
		compare  generic.Comparator[int]
		wantSize int
	}{
		{
			name:     "no err is ok",
			compare:  generic.ComparatorOrdered[int],
			wantSize: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantSize, sl.Len())
		})
	}
}

func TestSkipList_Cap(t *testing.T) {
	testCases := []struct {
		name     string
		compare  generic.Comparator[int]
		wantSize int
	}{
		{
			name:     "no err is ok",
			compare:  generic.ComparatorOrdered[int],
			wantSize: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantSize, sl.Cap())
		})
	}
}

func TestSkipList_AsSlice(t *testing.T) {
	testCases := []struct {
		name      string
		compare   generic.Comparator[int]
		wantSlice []int
	}{
		{
			name:      "no err is ok",
			compare:   generic.ComparatorOrdered[int],
			wantSlice: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
		})
	}
}

func ExampleNewSkipList() {
	l := list.NewSkipList[int](generic.ComparatorOrdered[int])
	l.Insert(123)
	val, _ := l.Get(0)
	fmt.Println(val)
	// output:
	// 123
}
