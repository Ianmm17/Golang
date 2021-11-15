package interview_problems_2

type Car struct {
	Make  string
	Model string
	Year  int
	OwnerIds []int
}

type Person struct {
	Id   int
	FirstName string
	LastName string
}

func GetCars() []Car {
	return []Car{
		{Make: "Ford", Model: "Model-T", Year: 1904, OwnerIds: []int{100, 400}},
		{Make: "Chevy", Model: "Cavalier", Year: 1994, OwnerIds: [] int{500}},
		{Make: "Toyota", Model: "Tundra", Year: 2016, OwnerIds: [] int{400}},
		{Make: "Honda", Model: "Accord", Year: 1998, OwnerIds: [] int{100}},
		{Make: "Mazda", Model: "Miata", Year: 1999, OwnerIds: [] int{200, 400}},
		{Make: "Nissan", Model: "Pathfinder", Year: 2018, OwnerIds: [] int{500}},
		{Make: "Ford", Model: "Fiesta", Year: 2002, OwnerIds: [] int{}},
		{Make: "Hyundai", Model: "Accent", Year: 2000, OwnerIds: [] int{500}},
		{Make: "Dodge", Model: "Durango", Year: 2005, OwnerIds: [] int{300}},
		{Make: "Lincoln", Model: "Town Car", Year: 1997, OwnerIds: [] int{200}},
	}
}

func GetPeople() []Person {
	return []Person{
		{Id: 100, FirstName: "Fred", LastName: "Flintstone"},
		{Id: 200, FirstName: "George", LastName: "Jetson"},
		{Id: 300, FirstName: "Bart", LastName: "Simpson"},
		{Id: 400, FirstName: "Bugs", LastName: "Bunny"},
		{Id: 500, FirstName: "Leroy", LastName: "Brown"},
	}
}
