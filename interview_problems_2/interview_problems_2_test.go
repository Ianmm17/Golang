package interview_problems_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestGetPersonsCars(t *testing.T) {
	t.Run("Should get mapped Fred's cars", func(t *testing.T) {
		// arrange
		fredId := 100
		cars := []string{"1904 Ford Model-T", "1998 Honda Accord"}
		expected := MappedCars{fredId, "Fred Flintstone", cars}

		// act
		actual := GetPeoplesCars(fredId)

		// assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Should get mapped George's cars", func(t *testing.T) {
		// arrange
		georgeId := 200
		cars := []string{"1999 Mazda Miata", "1997 Lincoln Town Car"}
		expected := MappedCars{georgeId, "George Jetson", cars}

		// act
		actual := GetPeoplesCars(georgeId)

		//assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Should get mapped Bart's cars", func(t *testing.T) {
		// arrange
		bartId := 300
		cars := []string{"2005 Dodge Durango"}
		expected := MappedCars{bartId, "Bart Simpson", cars}

		// act
		actual := GetPeoplesCars(bartId)

		//assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Should get mapped Bugs' cars", func(t *testing.T) {
		// arrange
		bugsId := 400
		cars := []string{"1904 Ford Model-T", "2016 Toyota Tundra", "1999 Mazda Miata"}
		expected := MappedCars{bugsId, "Bugs Bunny", cars}

		// act
		actual := GetPeoplesCars(bugsId)

		//assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Should get mapped Leroy's cars", func(t *testing.T) {
		// arrange
		leroyId := 500
		cars := []string{"1994 Chevy Cavalier", "2018 Nissan Pathfinder", "2000 Hyundai Accent"}
		expected := MappedCars{leroyId, "Leroy Brown", cars}

		// act
		actual := GetPeoplesCars(leroyId)

		//assert
		assert.Equal(t, expected, actual)
	})
}
