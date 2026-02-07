package slices_test

import (
	"github.com/synesissoftware/ANGoLS/slices"

	"github.com/stretchr/testify/assert"

	"testing"
)

func Benchmark_SelectSliceOfInt(b *testing.B) {

	inputs := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	for b.Loop() {
		for _, input := range inputs {
			slices.SelectSliceOfInt(input, func(_, _ int) (bool, error) {
				return true, nil
			})
		}
	}
}

func Benchmark_SelectSliceOfString(b *testing.B) {

	inputs := [][]string{
		{},
		{"1"},
		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	}

	for b.Loop() {
		for _, input := range inputs {
			slices.SelectSliceOfString(input, func(_ int, _ string) (bool, error) {
				return true, nil
			})
		}
	}
}

func Test_SelectSliceOfInt(t *testing.T) {

	{
		input := []int{}

		{
			r, _ := slices.SelectSliceOfInt(
				input,
				func(_, _ int) (bool, error) { return true, nil },
			)

			assert.Equal(t, []int{}, r)
		}
		{
			r, _ := slices.SelectSliceOfInt(
				input,
				func(_, _ int) (bool, error) { return false, nil },
			)

			assert.Equal(t, []int{}, r)
		}
	}

	{
		input := []int{1}

		{
			r, _ := slices.SelectSliceOfInt(
				input,
				func(_, _ int) (bool, error) { return true, nil },
			)

			assert.Equal(t, []int{1}, r)
		}
		{
			r, _ := slices.SelectSliceOfInt(
				input,
				func(_, _ int) (bool, error) { return false, nil },
			)

			assert.Equal(t, []int{}, r)
		}
	}

	{
		input := []int{1, 2, 3, 4, 5, 6}

		{
			r, _ := slices.SelectSliceOfInt(
				input,
				func(_, input_item int) (bool, error) { return 0 != (input_item % 2), nil },
			)

			assert.Equal(t, []int{1, 3, 5}, r)
		}
		{
			r, _ := slices.SelectSliceOfInt(
				input,
				func(_, input_item int) (bool, error) { return 0 == (input_item % 2), nil },
			)

			assert.Equal(t, []int{2, 4, 6}, r)
		}
	}
}

func Test_SelectSliceOfUint(t *testing.T) {

	{
		input := []uint{}

		{
			r, _ := slices.SelectSliceOfUint(
				input,
				func(_ int, _ uint) (bool, error) { return true, nil },
			)

			assert.Equal(t, []uint{}, r)
		}
		{
			r, _ := slices.SelectSliceOfUint(
				input,
				func(_ int, _ uint) (bool, error) { return false, nil },
			)

			assert.Equal(t, []uint{}, r)
		}
	}

	{
		input := []uint{1}

		{
			r, _ := slices.SelectSliceOfUint(
				input,
				func(_ int, _ uint) (bool, error) { return true, nil },
			)

			assert.Equal(t, []uint{1}, r)
		}
		{
			r, _ := slices.SelectSliceOfUint(
				input,
				func(_ int, _ uint) (bool, error) { return false, nil },
			)

			assert.Equal(t, []uint{}, r)
		}
	}

	{
		input := []uint{1, 2, 3, 4, 5, 6}

		{
			r, _ := slices.SelectSliceOfUint(
				input,
				func(_ int, input_item uint) (bool, error) { return 0 != (input_item % 2), nil },
			)

			assert.Equal(t, []uint{1, 3, 5}, r)
		}
		{
			r, _ := slices.SelectSliceOfUint(
				input,
				func(_ int, input_item uint) (bool, error) { return 0 == (input_item % 2), nil },
			)

			assert.Equal(t, []uint{2, 4, 6}, r)
		}
	}
}

func Test_SelectSliceOfInteger(t *testing.T) {

	// int
	{
		{
			input := []int{}

			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_, _ int) (bool, error) { return true, nil },
				)

				assert.Equal(t, []int{}, r)
			}
			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_, _ int) (bool, error) { return false, nil },
				)

				assert.Equal(t, []int{}, r)
			}
		}

		{
			input := []int{1}

			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_, _ int) (bool, error) { return true, nil },
				)

				assert.Equal(t, []int{1}, r)
			}
			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_, _ int) (bool, error) { return false, nil },
				)

				assert.Equal(t, []int{}, r)
			}
		}

		{
			input := []int{1, 2, 3, 4, 5, 6}

			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_, input_item int) (bool, error) { return 0 != (input_item % 2), nil },
				)

				assert.Equal(t, []int{1, 3, 5}, r)
			}
			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_, input_item int) (bool, error) { return 0 == (input_item % 2), nil },
				)

				assert.Equal(t, []int{2, 4, 6}, r)
			}
		}
	}

	// uint64
	{
		{
			input := []uint64{}

			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_ int, _ uint64) (bool, error) { return true, nil },
				)

				assert.Equal(t, []uint64{}, r)
			}
			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_ int, _ uint64) (bool, error) { return false, nil },
				)

				assert.Equal(t, []uint64{}, r)
			}
		}

		{
			input := []uint64{1}

			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_ int, _ uint64) (bool, error) { return true, nil },
				)

				assert.Equal(t, []uint64{1}, r)
			}
			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_ int, _ uint64) (bool, error) { return false, nil },
				)

				assert.Equal(t, []uint64{}, r)
			}
		}

		{
			input := []uint64{1, 2, 3, 4, 5, 6}

			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_ int, input_item uint64) (bool, error) { return 0 != (input_item % 2), nil },
				)

				assert.Equal(t, []uint64{1, 3, 5}, r)
			}
			{
				r, _ := slices.SelectSliceOfInteger(
					input,
					func(_ int, input_item uint64) (bool, error) { return 0 == (input_item % 2), nil },
				)

				assert.Equal(t, []uint64{2, 4, 6}, r)
			}
		}
	}
}

func Test_SelectSliceOfString(t *testing.T) {

	{
		input := []string{}

		{
			r, _ := slices.SelectSliceOfString(
				input,
				func(_ int, _ string) (bool, error) { return true, nil },
			)

			assert.Equal(t, []string{}, r)
		}
		{
			r, _ := slices.SelectSliceOfString(
				input,
				func(_ int, _ string) (bool, error) { return false, nil },
			)

			assert.Equal(t, []string{}, r)
		}
	}

	{
		input := []string{"a"}

		{
			r, _ := slices.SelectSliceOfString(
				input,
				func(_ int, _ string) (bool, error) { return true, nil },
			)

			assert.Equal(t, []string{"a"}, r)
		}
		{
			r, _ := slices.SelectSliceOfString(
				input,
				func(_ int, _ string) (bool, error) { return false, nil },
			)

			assert.Equal(t, []string{}, r)
		}
	}

	{
		input := []string{"", "a", "bcd", "d", "efghij", "kl"}

		{
			r, _ := slices.SelectSliceOfString(
				input,
				func(_ int, input_item string) (bool, error) { return len(input_item) < 2, nil },
			)

			assert.Equal(t, []string{"", "a", "d"}, r)
		}
		{
			r, _ := slices.SelectSliceOfString(
				input,
				func(_ int, input_item string) (bool, error) { return len(input_item) >= 2, nil },
			)

			assert.Equal(t, []string{"bcd", "efghij", "kl"}, r)
		}
	}
}
