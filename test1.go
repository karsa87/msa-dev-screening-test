package main

import (
	"fmt"
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	convertText := "1AB23C5678D"

	response := []string{}
	responseArrNumeric := []int{}
	responseArrAlphabet := []string{}
	for _, char := range convertText {
		response = append(response, string(char))

		if isNumeric(string(char)) == true {
			charInt, _ := strconv.Atoi(string(char))
			responseArrNumeric = append(responseArrNumeric, charInt)
		} else {
			responseArrAlphabet = append(responseArrAlphabet, string(char))
		}

	}

	fmt.Print("Result : ")
	fmt.Println(response)
	fmt.Print("Result : ")
	fmt.Println(responseArrNumeric)
	fmt.Print("Result : ")
	fmt.Println(responseArrAlphabet)
}
