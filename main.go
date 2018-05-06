package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const length = 16

func main() {
	fmt.Println(generate(""))
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

	return splitter(result.String())
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

func splitter(card string) string {
	var result bytes.Buffer
	for i := 0; i < len(card); i++ {
		if i > 0 && (i%4) == 0 {
			result.WriteString("-")
		}
		result.WriteString(string(card[i]))
	}
	return result.String()
}
