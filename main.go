package main

import (
	"fmt"
	//"practiceApp/practice"
	"practiceApp/goProblems"
)

func main() {
	choice := "add"
	doubtInt := goProblems.DoubleInt(2)
	doubleCollectionResults := goProblems.CollectionSliceInt([]int{5,6,7,})
	doWeMatch := goProblems.DoWeMatch([]int{1,2,3}, []int{1,2,3})
	lowHighResults := goProblems.LowestAndHighest([]int{3,4,5}, choice)
	numberOfLegs := goProblems.NumberOfLegs(2, 2, 1)
	sumInt := goProblems.SumInt([]int{1,2,3,4})
	sumIntMulti := goProblems.ProductMulti([]int{1,2,3,4})
	userChoice := goProblems.UserChoiceAddOrMultiply([]int{1,2,3,4}, false)
	convertToSeconds := goProblems.ConvertToSeconds(1, 2)
	minToSec := goProblems.MinToSec(4)
	hrToSec := goProblems.HrToSec(1)

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
	fmt.Println(name, ", Would you like to run Go problems? y/n")

	fmt.Println(doubtInt)
	fmt.Println(doubleCollectionResults)
	fmt.Println(doWeMatch)
	fmt.Printf("%v is the lowest and %v is the highest", lowHighResults[0], lowHighResults[1])
	fmt.Printf("\nYou have chosen %v and the total is: %v", choice, lowHighResults[2])
	fmt.Printf("\nThe total number of legs on your farm is: %v\n", numberOfLegs)
	fmt.Println(sumInt)
	fmt.Println(sumIntMulti)
	fmt.Println(userChoice)
	fmt.Println(convertToSeconds)
	fmt.Println(minToSec)
	fmt.Println(hrToSec)

}
