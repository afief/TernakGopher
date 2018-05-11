package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const length = 16

func main() {
	visa := flag.Bool("visa", false, "Generate Visa Credit Card Number")
	master := flag.Bool("master", true, "Generate Master Creedit Card Number")
	toValidate := flag.String("validate", "", "Validate existing credit card number")
	flag.Parse()

	generated := ""
	if *toValidate != "" {
		fmt.Println(validate(*toValidate))
	} else if *visa {
		generated = generate("4")
	} else if *master {
		generated = generate("54")
	}

	fmt.Println(generated)
}

func generate(bin string) string {
	randomLength := length - (len(bin) + 1)

	rand.Seed(time.Now().UnixNano())

	var result bytes.Buffer
	result.WriteString(bin)
	for i := 0; i < randomLength; i++ {
		result.WriteString(strconv.Itoa(rand.Intn(10)))
	}

	result.WriteString(getLuhnChecksum(result.String()))

	return addDash(result.String())
}

func getLuhnChecksum(card string) string {
	sum := 0
	lenOffset := (length + 1) % 2
	var t int
	for pos := 0; pos < (length - 1); pos++ {
		if ((pos + lenOffset) % 2) > 0 {
			t, _ = strconv.Atoi(string(card[pos]))
			t *= 2
			if t > 9 {
				t -= 9
			}
			sum += t
		} else {
			t, _ = strconv.Atoi(string(card[pos]))
			sum += t
		}
	}
	return strconv.Itoa((10 - (sum % 10)) % 10)
}

func addDash(card string) string {
	var result bytes.Buffer
	for i := 0; i < len(card); i++ {
		if i > 0 && (i%4) == 0 {
			result.WriteString("-")
		}
		result.WriteString(string(card[i]))
	}
	return result.String()
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

func validate(dashedCard string) bool {
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
