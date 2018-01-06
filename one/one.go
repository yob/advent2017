package one

import (
	"strconv"
	"strings"
)

func stringToIntArray(value string) []int {
	characters := strings.Split(value, "")
	numbers := make([]int, len(characters))
	for i, c := range characters {
		numbers[i], _ = strconv.Atoi(c)
	}
	return numbers
}

func Sum(value string) int {
	numbers := stringToIntArray(value)
	sum := 0
	for i, _ := range numbers {
		compare_i := (i+1) % len(numbers);
		if numbers[i] == numbers[compare_i] {
			sum += numbers[i]
		}
	}
	return sum
}

func SumIfOpposite(value string) int {
	numbers := stringToIntArray(value)
	sum := 0
	for i, _ := range numbers {
		compare_i := (i+(len(numbers)/2)) % len(numbers);
		if numbers[i] == numbers[compare_i] {
			sum += numbers[i]
		}
	}
	return sum
}
