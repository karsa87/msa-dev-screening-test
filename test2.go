package main

import (
	"fmt"
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func calculate(num1 int, num2 int, operator string) int {
	switch string(operator) {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case ":":
		return num1 / num2
	default:
		return 0
	}
}

func main() {
	convertText := "5+5+5*5:5"

	response := 0
	operator := ""
	for _, char := range convertText {
		if isNumeric(string(char)) {
			if operator != "" {
				charInt, _ := strconv.Atoi(string(char))
				response = calculate(response, charInt, operator)
			} else {
				charInt, _ := strconv.Atoi(string(char))
				response = charInt
			}
		} else {
			operator = string(char)
		}
	}

	fmt.Println(response)
	if response == 15 {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}
}
