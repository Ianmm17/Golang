package main

import (
	"fmt"
	"practiceApp/goProblems"
)

func main() {

	userName := getUserName()

	doYouWantToRun(userName)

}

func doYouWantToRun(name string) string {
	var yesOrNo string
	fmt.Println(name, ", Would you like to run Go problems? y/n")
	fmt.Scanf("%s", &yesOrNo)

	if yesOrNo == "y" {
		goProblems.ListOfGoProblems()
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
