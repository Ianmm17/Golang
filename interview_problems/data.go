package interview_problems

type Car struct {
	Id    int
	Make  string
	Model string
	Year  int
}

type Person struct {
	Id   int
	Name string
	Cars []int
}

func GetCars() []Car {
	return []Car{
		{Id: 100, Make: "Ford", Model: "Model-T", Year: 1904},
		{Id: 200, Make: "Chevy", Model: "Cavalier", Year: 1994},
		{Id: 300, Make: "Toyota", Model: "Tundra", Year: 2016},
		{Id: 400, Make: "Honda", Model: "Accord", Year: 1998},
		{Id: 500, Make: "Mazda", Model: "Miata", Year: 1999},
		{Id: 600, Make: "Nissan", Model: "Pathfinder", Year: 2018},
		{Id: 700, Make: "Ford", Model: "Fiesta", Year: 2002},
		{Id: 800, Make: "Hyundai", Model: "Accent", Year: 2000},
		{Id: 900, Make: "Dodge", Model: "Durango", Year: 2005},
		{Id: 1000, Make: "Lincoln", Model: "Town Car", Year: 1997},
	}
}

func GetPeople() []Person {
	return []Person{
		{Id: 100, Name: "Fred Flintstone", Cars: []int{400, 100}},
		{Id: 200, Name: "George Jetson", Cars: []int{1000, 500}},
		{Id: 300, Name: "Bart Simpson", Cars: []int{900}},
		{Id: 400, Name: "Bugs Bunny", Cars: []int{300, 100, 500}},
		{Id: 500, Name: "Leroy Brown", Cars: []int{600, 200, 800}},
	}
}
