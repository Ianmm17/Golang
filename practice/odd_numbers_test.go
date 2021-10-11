package practice_test

import (
	"practiceApp/practice"
	"testing"
)


func TestAddOddNumbers1(t *testing.T) {
	//setup
	tests := []struct {
		start, end, expect int
		wantErr bool
	}{
		{start: 1, end: 10, expect: 25},
		{start: 1, end: 5, expect: 9},
		{start: 2, end: 5, expect: 8},
		{start: 2, end: 3, expect: 3},
		{start: 2, end: 2, expect: 0},
		{start: 11, end: 1, wantErr: true},
		{start: -10, end: -1, expect: -25},
		{start: -10, end: 10, expect: 0},
		{start:  1, end:    100, expect: 2500},
	}
	for _, testCase := range tests {
		t.Run("Starting", func(t *testing.T) {
			//act
			actual, err := practice.AddOddNumbers(testCase.start, testCase.end)
			//assert
			if err != nil && !testCase.wantErr {
				t.Errorf("unexpected error %v", err)
			}else if err == nil && testCase.wantErr {
				t.Errorf("expected error, got no error")
			}else if actual != testCase.expect {
				t.Errorf("AddOddNumbers() = %v, expect %v", actual, testCase.expect)
			}
		})
	}
}
