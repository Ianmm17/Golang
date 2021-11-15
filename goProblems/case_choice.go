package goProblems

import (
	"fmt"
	"strconv"
	"strings"
)

func ListOfGoProblems() {
	var numChoice string
	println("1. doubles integer")
	println("2. doubles each integer in given slice")

	fmt.Scanf("%s", &numChoice)
	fmt.Println("You have chosen: ", numChoice)

	switch numChoice {
	case "1":
		Case1()
	case "2":
		Case2()
	default:
		fmt.Println(numChoice, "isn't an option. Please choose from the following.")
		ListOfGoProblems()
	}

}

func Case1() {
	var passedInt int
	fmt.Println("Enter Number: ")
	if _, err := fmt.Scan(&passedInt); err != nil {
		println("Integer not given, please enter integer")
		Case1()
	} else {
		fmt.Println("Your doubled integer is: ", DoubleInt(passedInt))
	}
}

//TODO: look into removing white space in giving slice
func Case2() error {
	var passedStr string
	fmt.Println("Enter Numbers: ")
	fmt.Scan(&passedStr)
	strSlice := strings.Split(passedStr, ",")
	convertedUsers, err := convertUserInput(strSlice, Case2)
	if err != nil {
		return err
	}

	fmt.Println("Your doubled integers are: ", CollectionSliceInt(convertedUsers))

	ListOfGoProblems()

	return nil
}

func Case3() {
	var passedInt int
	fmt.Println("Enter Number: ")
	if _, err := fmt.Scan(&passedInt); err != nil {
		println("Integer not given, please enter integer")
		Case1()
	} else {
		fmt.Println("Your doubled integer is: ", DoubleInt(passedInt))
	}
}

func Case4() {
	var passedInt int
	fmt.Println("Enter Number: ")
	if _, err := fmt.Scan(&passedInt); err != nil {
		println("Integer not given, please enter integer")
		Case1()
	} else {
		fmt.Println("Your doubled integer is: ", DoubleInt(passedInt))
	}
}

func Case5() {
	var passedInt int
	fmt.Println("Enter Number: ")
	if _, err := fmt.Scan(&passedInt); err != nil {
		println("Integer not given, please enter integer")
		Case1()
	} else {
		fmt.Println("Your doubled integer is: ", DoubleInt(passedInt))
	}
}

func convertUserInput(strSlice[]string, reCallCase func() error) ([]int, error) {
	var allInOneSlice []int
	for _, v := range strSlice {
		converted, err := strconv.Atoi(v)
		if err != nil {
			println("Input error, please insert integers seperated by commas")
			reCallCase()
			return nil, err
		}
		allInOneSlice = append(allInOneSlice, converted)
	}
	return allInOneSlice, nil
}
