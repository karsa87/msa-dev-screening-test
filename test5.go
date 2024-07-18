package main

import (
	"fmt"
)

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func visualization(numbers []int) {
	max := findMax(numbers)

	for max >= 0 {
		for i := 0; i < len(numbers); i++ {
			if max == 0 {
				fmt.Print(numbers[i])
			} else if numbers[i] >= max {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		max--
		fmt.Println("")
	}
}

func main() {
	numbers := []int{1, 4, 5, 6, 8, 2}

	ok := 1
	temp := 0
	i := len(numbers) - 1
	for ok <= len(numbers) {
		newIndex := i
		for j := len(numbers) - 1; j >= 0; j-- {
			if numbers[newIndex] < numbers[j] {
				temp = numbers[j]
				numbers[j] = numbers[newIndex]
				numbers[newIndex] = temp

				newIndex = j
				visualization(numbers)
			}
		}
		ok++
	}
}
