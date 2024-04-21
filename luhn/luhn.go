package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func sumDigits(digits string) (int, error) {
	var sum int = 0
	for i := 0; i < len(digits); i++ {
		val, err := strconv.Atoi(digits[i : i+1])
		if err != nil {
			return 0, err
		}
		sum += val
	}
	return sum, nil
}

func calculateCheckDigit(payload string) (int, error) {
	var digitsBuilder strings.Builder
	var useDouble bool = true

	for i := len(payload) - 1; i >= 0; i-- {
		item, err := strconv.Atoi(payload[i : i+1])
		if err != nil {
			return 0, err
		}
		if useDouble {
			item *= 2
		}
		digitsBuilder.WriteString(fmt.Sprint(item))
		useDouble = !useDouble
	}

	var digits string = digitsBuilder.String()
	sumDigits, err := sumDigits(digits)
	if err != nil {
		return 0, err
	}

	var checkDigit int = 10 - (sumDigits % 10)
	return checkDigit, nil
}

func Validate(number string) bool {
	var checkDigit string = number[len(number)-1:]
	var payload string = number[:len(number)-1]
	calculatedCheckDigit, err := calculateCheckDigit(payload)
	if err != nil {
		return false
	}
	return fmt.Sprint(calculatedCheckDigit) == checkDigit
}
