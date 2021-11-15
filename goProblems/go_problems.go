package goProblems

func DoubleInt(num int) int {
	return num * 2
}

func CollectionSliceInt(nums []int) []int {
	var newSlice []int
	for i := 0; i < len(nums); i++ {
		double := nums[i] * 2
		newSlice = append(newSlice, double)
	}
	return newSlice
}

func DoWeMatch(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func LowestAndHighest(nums []int, choice string) []int {

	if len(nums) == 0 {
		return []int{}
	}

	totalAdd := 0
	low := nums[0]
	high := nums[0]


	for i := 1; i < len(nums); i++ {
		if nums[i] > high {
			high = nums[i]
		}
		if nums[i] < low {
			low = nums[i]
		}
	}
	if choice == "add" {
		totalAdd = low + high
	}
	if choice == "multiply" {
		totalAdd = low * high
	}

	return []int{low, high, totalAdd}
}

//NumberOfLegs TODO: make error out if it receives negative number
func NumberOfLegs(numSheep, numChickens, numOctopi int) int {
	numSheepLegs := numSheep * 4
	numChickenLegs := numChickens * 2
	numOctopiLegs := numOctopi * 8

	totalNumOfLegs := numSheepLegs + numChickenLegs + numOctopiLegs

	return totalNumOfLegs
}

func SumInt(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func ProductMulti(slice []int) int {
	sum := 1
	for i := 0; i < len(slice); i++ {
		sum *= slice[i]
	}
	return sum
}

func UserChoiceAddOrMultiply(slice []int, choice bool) int {
	if choice {
		return SumInt(slice)
	} else {
		return ProductMulti(slice)
	}
}

func ConvertToSeconds(hour int, min int) int {
	return HrToSec(hour) + MinToSec(min)
}

func MinToSec(min int) int {
	return min * 60
}

func HrToSec(hr int) int {
	return hr * 3600
}

func CalcFactorial(num int) int {
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	return factorial
}

func RecursiveFactorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * RecursiveFactorial(num-1)
	}

}
