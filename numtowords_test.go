package numtowords_test

import (
	"testing"

	"github.com/Sameeksha21/numtowords"
)

func TestInvalidInput(t *testing.T) {
	_, err := numtowords.Convert(numtowords.MaxNum + 1)
	if err == nil {
		t.Log("Expected error for input greater than MaxNum, got nil")
		t.Fail()
	}
	_, err = numtowords.Convert(numtowords.MinNum - 1)
	if err == nil {
		t.Log("Expected error for input less than MinNum, got nil")
		t.Fail()
	}
}

func TestZeroInput(t *testing.T) {
	result, err := numtowords.Convert(0)
	if err != nil {
		t.Log("Unexpected error for input 0")
		t.FailNow()
	}
	if result != "zero" {
		t.Logf("Expected 'zero' for input 0, got '%s'", result)
		t.Fail()
	}
}

func TestUnits(t *testing.T) {
	units := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	for i, v := range units {
		t.Logf("Testing for the number %v", i)
		result, err := numtowords.Convert(i)
		if err != nil {
			t.Logf("Unexpected error for input %d", i)
			t.Fail()
		}
		if result != units[i] {
			t.Logf("Expected '%s' for input %d, got '%s'", v, i, result)
			t.Fail()
		}
	}
}
func TestTens(t *testing.T) {
	testcases := map[int]string{
		20: "twenty",
		35: "thirty five",
		99: "ninety nine",
		56: "fifty six",
	}
	for k, v := range testcases {
		t.Logf("Testing for the number %v", k)
		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("Unexpected error for input %d", k)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%s' for input %d, got '%s'", v, k, result)
			t.Fail()
		}
	}
}

func TestHundreds(t *testing.T) {
	testCases := map[int]string{
		109: "one hundred and nine",
		340: "three hundred and forty",
		313: "three hundred and thirteen",
		700: "seven hundred",
	}
	for k, v := range testCases {
		t.Logf("Testing for the number %v", k)
		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("Unexpected error for input %d", k)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%s' for input %d, got '%s'", v, k, result)
			t.Fail()
		}
	}
}

func TestNegativeNumbers(t *testing.T) {
	testCases := map[int]string{
		0:    "zero",
		-1:   "minus one",
		-15:  "minus fifteen",
		-123: "minus one hundred and twenty three",
		-109: "minus one hundred and nine",
		-980: "minus nine hundred and eighty",
	}
	for k, v := range testCases {
		t.Logf("Testing for the number %v", k)
		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("Unexpected error for input %d", k)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%s' for input %d, got '%s'", v, k, result)
			t.Fail()
		}
	}
}
