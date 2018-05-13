package validator

import (
	"bytes"
	"strconv"
)

// Validate is to validate the existing credit card number
func Validate(dashedCard string) bool {
	card := removeDash(dashedCard)

	checksum := 0
	for i := 2 - (len(card) % 2); i <= len(card); i += 2 {
		nc, _ := strconv.Atoi(string(card[i-1]))
		checksum += nc
	}

	for i := (len(card) % 2) + 1; i < len(card); i += 2 {
		digit, _ := strconv.Atoi(string(card[i-1]))
		digit *= 2

		if digit < 10 {
			checksum += digit
		} else {
			checksum += digit - 9
		}
	}

	if (checksum % 10) == 0 {
		return true
	}
	return false
}

func removeDash(card string) string {
	var result bytes.Buffer
	for i := 0; i < len(card); i++ {
		if string(card[i]) != "-" {
			result.WriteString(string(card[i]))
		}
	}
	return result.String()
}
