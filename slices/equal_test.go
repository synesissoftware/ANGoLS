package slices_test

import (
	"github.com/synesissoftware/ANGoLS/slices"

	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_EqualSliceOfInt(t *testing.T) {

	{
		assert.True(t, slices.EqualSliceOfInt([]int{}, []int{}))
		assert.True(t, slices.EqualSliceOfInt([]int{1}, []int{1}))
		assert.False(t, slices.EqualSliceOfInt([]int{1}, []int{2}))

		assert.True(t, slices.EqualSliceOfInt([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInt([]int{1, 2, 3, 4}, []int{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInt([]int{1, 2, 3, 4, 5}[:4], []int{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]int{}, []int{}))
		assert.True(t, slices.EqualSliceOfInteger([]int{1}, []int{1}))
		assert.False(t, slices.EqualSliceOfInteger([]int{1}, []int{2}))

		assert.True(t, slices.EqualSliceOfInteger([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]int{1, 2, 3, 4}, []int{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]int{1, 2, 3, 4, 5}[:4], []int{1, 2, 3, 4, 6}[:4]))
	}
}

func Test_EqualSliceOfUint(t *testing.T) {

	{
		assert.True(t, slices.EqualSliceOfUint([]uint{}, []uint{}))
		assert.True(t, slices.EqualSliceOfUint([]uint{1}, []uint{1}))
		assert.False(t, slices.EqualSliceOfUint([]uint{1}, []uint{2}))

		assert.True(t, slices.EqualSliceOfUint([]uint{1, 2, 3, 4}, []uint{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfUint([]uint{1, 2, 3, 4}, []uint{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfUint([]uint{1, 2, 3, 4, 5}[:4], []uint{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uint{}, []uint{}))
		assert.True(t, slices.EqualSliceOfInteger([]uint{1}, []uint{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uint{1}, []uint{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uint{1, 2, 3, 4}, []uint{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uint{1, 2, 3, 4}, []uint{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uint{1, 2, 3, 4, 5}[:4], []uint{1, 2, 3, 4, 6}[:4]))
	}
}

func Test_EqualSliceOfInteger(t *testing.T) {

	{
		assert.True(t, slices.EqualSliceOfInteger([]int{}, []int{}))
		assert.True(t, slices.EqualSliceOfInteger([]int{1}, []int{1}))
		assert.False(t, slices.EqualSliceOfInteger([]int{1}, []int{2}))

		assert.True(t, slices.EqualSliceOfInteger([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]int{1, 2, 3, 4}, []int{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]int{1, 2, 3, 4, 5}[:4], []int{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uint{}, []uint{}))
		assert.True(t, slices.EqualSliceOfInteger([]uint{1}, []uint{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uint{1}, []uint{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uint{1, 2, 3, 4}, []uint{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uint{1, 2, 3, 4}, []uint{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uint{1, 2, 3, 4, 5}[:4], []uint{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]int8{}, []int8{}))
		assert.True(t, slices.EqualSliceOfInteger([]int8{1}, []int8{1}))
		assert.False(t, slices.EqualSliceOfInteger([]int8{1}, []int8{2}))

		assert.True(t, slices.EqualSliceOfInteger([]int8{1, 2, 3, 4}, []int8{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]int8{1, 2, 3, 4}, []int8{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]int8{1, 2, 3, 4, 5}[:4], []int8{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uint8{}, []uint8{}))
		assert.True(t, slices.EqualSliceOfInteger([]uint8{1}, []uint8{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uint8{1}, []uint8{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uint8{1, 2, 3, 4}, []uint8{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uint8{1, 2, 3, 4}, []uint8{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uint8{1, 2, 3, 4, 5}[:4], []uint8{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]int16{}, []int16{}))
		assert.True(t, slices.EqualSliceOfInteger([]int16{1}, []int16{1}))
		assert.False(t, slices.EqualSliceOfInteger([]int16{1}, []int16{2}))

		assert.True(t, slices.EqualSliceOfInteger([]int16{1, 2, 3, 4}, []int16{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]int16{1, 2, 3, 4}, []int16{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]int16{1, 2, 3, 4, 5}[:4], []int16{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uint16{}, []uint16{}))
		assert.True(t, slices.EqualSliceOfInteger([]uint16{1}, []uint16{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uint16{1}, []uint16{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uint16{1, 2, 3, 4}, []uint16{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uint16{1, 2, 3, 4}, []uint16{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uint16{1, 2, 3, 4, 5}[:4], []uint16{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]int32{}, []int32{}))
		assert.True(t, slices.EqualSliceOfInteger([]int32{1}, []int32{1}))
		assert.False(t, slices.EqualSliceOfInteger([]int32{1}, []int32{2}))

		assert.True(t, slices.EqualSliceOfInteger([]int32{1, 2, 3, 4}, []int32{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]int32{1, 2, 3, 4}, []int32{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]int32{1, 2, 3, 4, 5}[:4], []int32{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uint32{}, []uint32{}))
		assert.True(t, slices.EqualSliceOfInteger([]uint32{1}, []uint32{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uint32{1}, []uint32{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uint32{1, 2, 3, 4}, []uint32{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uint32{1, 2, 3, 4}, []uint32{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uint32{1, 2, 3, 4, 5}[:4], []uint32{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]int64{}, []int64{}))
		assert.True(t, slices.EqualSliceOfInteger([]int64{1}, []int64{1}))
		assert.False(t, slices.EqualSliceOfInteger([]int64{1}, []int64{2}))

		assert.True(t, slices.EqualSliceOfInteger([]int64{1, 2, 3, 4}, []int64{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]int64{1, 2, 3, 4}, []int64{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]int64{1, 2, 3, 4, 5}[:4], []int64{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uint64{}, []uint64{}))
		assert.True(t, slices.EqualSliceOfInteger([]uint64{1}, []uint64{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uint64{1}, []uint64{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uint64{1, 2, 3, 4}, []uint64{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uint64{1, 2, 3, 4}, []uint64{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uint64{1, 2, 3, 4, 5}[:4], []uint64{1, 2, 3, 4, 6}[:4]))
	}

	{
		assert.True(t, slices.EqualSliceOfInteger([]uintptr{}, []uintptr{}))
		assert.True(t, slices.EqualSliceOfInteger([]uintptr{1}, []uintptr{1}))
		assert.False(t, slices.EqualSliceOfInteger([]uintptr{1}, []uintptr{2}))

		assert.True(t, slices.EqualSliceOfInteger([]uintptr{1, 2, 3, 4}, []uintptr{1, 2, 3, 4}))
		assert.False(t, slices.EqualSliceOfInteger([]uintptr{1, 2, 3, 4}, []uintptr{1, 2, 4, 3}))

		assert.True(t, slices.EqualSliceOfInteger([]uintptr{1, 2, 3, 4, 5}[:4], []uintptr{1, 2, 3, 4, 6}[:4]))
	}
}

func Test_EqualSliceOfString(t *testing.T) {

	{
		assert.True(t, slices.EqualSliceOfString([]string{}, []string{}))
		assert.True(t, slices.EqualSliceOfString([]string{"1"}, []string{"1"}))
		assert.False(t, slices.EqualSliceOfString([]string{"1"}, []string{"2"}))

		assert.True(t, slices.EqualSliceOfString([]string{"1", "2", "3", "4"}, []string{"1", "2", "3", "4"}))
		assert.False(t, slices.EqualSliceOfString([]string{"1", "2", "3", "4"}, []string{"1", "2", "4", "3"}))

		assert.True(t, slices.EqualSliceOfString([]string{"1", "2", "3", "4", "5"}[:4], []string{"1", "2", "3", "4", "6"}[:4]))
	}
}
