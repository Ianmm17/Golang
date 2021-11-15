package interview_problems_2

import "fmt"

type MappedCars struct {
	Id int
	Name string
	Cars []string
}

func GetPeoplesCars(personsId int) MappedCars {
	// var ownedCars []string
	var ownersCars []string
	// var ownedCarStr || formattedCar
	var ownersCarsStr string
	var personsName string

	for _, car := range GetCars() {
		// ownerId
		for _, id := range car.OwnerIds {
			if personsId == id {
				// You could do without this ownersCarsStr variable and inline it in the append
				ownersCarsStr = fmt.Sprintf("%d %s %s", car.Year, car.Make, car.Model)
				ownersCars = append(ownersCars, ownersCarsStr)
			}

		}
	}

	// name -> person
	for _, name := range GetPeople(){
		if personsId == name.Id {
			personsName = fmt.Sprintf("%s %s", name.FirstName, name.LastName)
		}
	}

	return MappedCars{
		Id: personsId,
		Name: personsName,
		Cars: ownersCars,
	}
}
