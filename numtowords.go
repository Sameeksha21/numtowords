package numtowords

import (
	"fmt"
	"math"
)

// MaxNum is the maximum number that can be converted to words
const MaxNum = 999

// MinNum is the minimum number that can be converted to words
const MinNum = -999

// Convert converts a specified number to words
func Convert(num int) (string, error) {
	if num > MaxNum || num < MinNum {
		return "", fmt.Errorf("Can only convert numbers between %v and %v", MinNum, MaxNum)
	}
	if num == 0 {
		return "zero", nil
	}
	units := [20]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := [8]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	result := ""
	if num < 0 {
		result += "minus "
		//num = -num
		num = int(math.Abs(float64(num)))
	}
	if num > 99 {
		hundredindex := num / 100
		result += units[hundredindex] + " hundred"
		num = num % 100
		if num == 0 {
			return result, nil
		}
		if num > 0 {
			result += " and "
		}
	}
	if num > 19 {
		{
			tenindex := num/10 - 2
			result += tens[tenindex]
			num = num % 10
		}
		if num == 0 {
			return result, nil

		}
		result += " "
	}
	result += units[num]
	return result, nil
}
