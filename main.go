package main

import (
	"fmt"
	"practiceApp/goProblems"
)

func main() {
	userName := getUserName()

	doYouWantToRun(userName)

}

func listOfGoProblems() {
	var numChoice string
	println("1. doubles integer")
	println("2. doubles each integer in given slice")

	fmt.Scanf("%s", &numChoice)
	fmt.Println("You have chosen: ", numChoice)

	switch numChoice {
	case "1":
		case1()
	case "2":
	default:
		fmt.Println(numChoice, "isn't an option. Please choose from the following.")
		listOfGoProblems()
	}

}

func doYouWantToRun(name string) string {
	var yesOrNo string
	fmt.Println(name, ", Would you like to run Go problems? y/n")
	fmt.Scanf("%s", &yesOrNo)

	if yesOrNo == "y" {
		listOfGoProblems()
	} else if yesOrNo == "n" {
		fmt.Println("no")
	} else {
		println("Invalid argument please try again!")
		doYouWantToRun(name)
	}
	return yesOrNo
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
	return name
}

func case1() {
	var passedInt int
	fmt.Println("Enter Number: ")
	if _, err := fmt.Scan(&passedInt); err != nil {
		println("Integer not given, please enter integer")
		case1()
	} else {
		fmt.Println("Your doubled integer is: ", goProblems.DoubleInt(passedInt))
	}
}
