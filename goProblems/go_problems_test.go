package goProblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleInt(t *testing.T) {

	t.Run("double the inputted number", func(t *testing.T) {
		num := 1
		//setup
		expected := 2
		//act
		actual := DoubleInt(num)
		//assert
		assert.Equal(t, expected, actual)
	})
	t.Run("handle negative input", func(t *testing.T) {
		num := -1
		expected := -2

		actual := DoubleInt(num)

		assert.Equal(t, expected, actual)
	})
	t.Run("handle when inputted number is zero", func(t *testing.T) {
		num := 0
		expected := 0

		actual := DoubleInt(num)

		assert.Equal(t, expected, actual)
	})
}

func TestCollectionSliceInt(t *testing.T) {


	t.Run("double each value in slice", func(t *testing.T) {
		nums := []int{1,2,3}
		expected := []int{2,4,6}

		actual := CollectionSliceInt(nums)

		assert.Equal(t, expected, actual)
	})
	/*t.Run("handles empty slice", func(t *testing.T) {
		var nums []int

		expected := []int{}
		fmt.Println(expected)
		actual := CollectionSliceInt(nums)

		assert.Equal(t, expected, actual)
	})*/
}

func TestDoWeMatch(t *testing.T) {

	t.Run("returns true for matching slices", func(t *testing.T) {
		slice1 := []int{1,2,3}
		slice2 := []int{1,2,3}

		expected := true

		actual := DoWeMatch(slice1, slice2)

		assert.Equal(t, expected, actual)
	})
	t.Run("returns false for mismatched slices", func(t *testing.T) {
		slice1 := []int{1,2,3}
		slice2 := []int{1,2,2}

		expected := false

		actual := DoWeMatch(slice1, slice2)

		assert.Equal(t, expected, actual)
	})
	t.Run("different lengths of slices returns false", func(t *testing.T) {
		slice1 := []int{1,2,3,4}
		slice2 := []int{1,2,2}

		expected := false

		actual := DoWeMatch(slice1, slice2)

		assert.Equal(t, expected, actual)
	})
}

func TestLowestAndHighest(t *testing.T) {

	t.Run("handles adding lowest and highest values", func(t *testing.T) {
		nums := []int{1,4,5}
		choice := "add"
		expected := []int{1,5,6}

		actual := LowestAndHighest(nums, choice)

		assert.Equal(t, expected, actual)
	})
	t.Run("handles multiplying lowest and highest values", func(t *testing.T) {
		nums := []int{1,4,5}
		choice := "multiply"
		expected := []int{1,5,5}

		actual := LowestAndHighest(nums, choice)

		assert.Equal(t, expected, actual)
	})
	t.Run("when slice is empty", func(t *testing.T) {
		nums := []int{}
		choice := "add"
		expected := []int{}

		actual := LowestAndHighest(nums, choice)

		assert.Equal(t, expected, actual)
	})
}

func TestNumberOfLegs(t *testing.T) {

	t.Run("returns total number of legs", func(t *testing.T) {
		numOfSheeps := 1
		numOfChickens := 1
		numOfOctopus := 1

		expected := 14

		actual := NumberOfLegs(numOfSheeps, numOfChickens, numOfOctopus)

		assert.Equal(t, expected, actual)
	})
}

func TestSumInt(t *testing.T) {

	t.Run("adds all numbers in slice and returns sum", func(t *testing.T) {
		slice := []int{1,2}

		expected := 3

		actual := SumInt(slice)

		assert.Equal(t, expected, actual)
	})
	t.Run("handles adding negative numbers", func(t *testing.T) {
		slice := []int{1,2,-4}

		expected := -1

		actual := SumInt(slice)

		assert.Equal(t, expected, actual)
	})
}

func TestProductMulti(t *testing.T) {

	t.Run("multiply all numbers in slice and returns product", func(t *testing.T) {
		slice := []int{1,2}

		expected := 2

		actual := ProductMulti(slice)

		assert.Equal(t, expected, actual)
	})
	t.Run("handles multiplying negative numbers", func(t *testing.T) {
		slice := []int{1,-2}

		expected := -2

		actual := ProductMulti(slice)

		assert.Equal(t, expected, actual)
	})
}

func TestUserChoiceAddOrMultiply(t *testing.T) {

	t.Run("when choice is true return sum", func(t *testing.T) {
		slice := []int{1,2}
		choice := true

		expected := 3

		actual := UserChoiceAddOrMultiply(slice, choice)

		assert.Equal(t, expected, actual)
	})
	t.Run("when choice is false return product", func(t *testing.T) {
		slice := []int{1,2}
		choice := false

		expected := 2

		actual := UserChoiceAddOrMultiply(slice, choice)

		assert.Equal(t, expected, actual)
	})
}

func TestConvertToSeconds(t *testing.T) {

	t.Run("convert hours and minutes to seconds", func(t *testing.T) {
		hour, min := 1,1

		expected := 3660

		actual := ConvertToSeconds(hour, min)

		assert.Equal(t, expected, actual)
	})
}
