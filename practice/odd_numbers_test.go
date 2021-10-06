package practice

import "testing"


func TestAddOddNumbers1(t *testing.T) {
	//setup
	tests := []struct {
		start, end, expect int
	}{
		{start: 1, end: 10, expect: 25},
		{start: 1, end: 5, expect: 9},
		{
			start:  1,
			end:    100,
			expect: 2500,
		},
	}
	for _, testCase := range tests {
		t.Run("Starting", func(t *testing.T) {
			//act
			actual := AddOddNumbers(testCase.start, testCase.end)
			//assert
			if actual != testCase.expect {
				t.Errorf("AddOddNumbers() = %v, expect %v", actual, testCase.expect)
			}
		})
	}
}
