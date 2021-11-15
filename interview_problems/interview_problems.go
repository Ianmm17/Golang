package interview_problems

import (
	"fmt"
)

func Interview() {
	fmt.Println("Get People: ", GetPeople())
	fmt.Println("Get Cars: ", GetCars())
}

type Result struct {
	Id   int
	Name string
	Car  []string
}

func GetPeoplesCars(personsID int) Result {
	var OwnerName string
	var OwnersCars []int
	var OwnersCarsStr string
	var OwnersCarsStrSlice []string
	for _, personInfo := range GetPeople() {
		if personsID == personInfo.Id {
			OwnerName = personInfo.Name
			OwnersCars = personInfo.Cars
		}

	}
	for _, car := range GetCars() {
		for _, eachOwnerCars := range OwnersCars {
			if car.Id == eachOwnerCars {
				OwnersCarsStr = fmt.Sprintf("%d %s %s", car.Year, car.Make, car.Model)
				OwnersCarsStrSlice = append(OwnersCarsStrSlice, OwnersCarsStr)
			}
		}
	}

	return Result{
		Id:   personsID,
		Name: OwnerName,
		Car:  OwnersCarsStrSlice,
	}
}
