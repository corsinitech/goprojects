package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("An error has occured: %s", err)
	}
	num2, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Printf("An erro has occured: %s", err)
	}
	operator := os.Args[3]

    var result float64

	switch operator {
	case "+":
        result = addition(num1, num2)
	case "-":
        result = subtraction(num1, num2)
    case "x":
        result = multiplication(num1, num2)
    // TODO: Figure out why '*' doesn't actually switch and run the function...
    case "*":
        result = multiplication(num1, num2)
	case "/":
        result = division(num1, num2)
	case "%":
        result = modulus(num1, num2)
	}

    fmt.Println(result)

}

func addition(num1, num2 float64 ) float64 { 
    return num1 + num2
}

func subtraction(num1, num2 float64) float64 {
    return num1 - num2
}

func multiplication(num1, num2 float64) float64 {
    return num1 * num2
}

func division(num1, num2 float64) float64 {
    return num1 / num2
}

func modulus(num1, num2 float64) float64 {
    return math.Mod(num1, num2)
}
