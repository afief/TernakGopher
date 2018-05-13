package main

import (
	"flag"
	"fmt"

	"github.com/afief/ternakgopher/pkg/generator"
	"github.com/afief/ternakgopher/pkg/validator"
)

func main() {
	visa := flag.Bool("visa", false, "Generate Visa Credit Card Number")
	master := flag.Bool("master", true, "Generate Master Creedit Card Number")
	toValidate := flag.String("validate", "", "Validate existing credit card number")
	flag.Parse()

	generated := ""
	if *toValidate != "" {
		fmt.Println(validator.Validate(*toValidate))
	} else if *visa {
		generated = generator.Generate("4")
	} else if *master {
		generated = generator.Generate("54")
	}

	fmt.Println(generated)
}
