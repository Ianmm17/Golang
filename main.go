package main

import (
	"fmt"
	"practiceApp/practice"
)

func main() {
	fmt.Println("Go Practice!")
	addedOdd, err := practice.AddOddNumbers(222, 150)
	if err != nil {
		println(err.Error())
	} else {
		println(addedOdd)
	}
}
