package two

import (
	"strconv"
	"strings"
)

func stringToIntArray(value string) []int {
	characters := strings.Split(value, "\t")
	numbers := make([]int, len(characters))
	for i, c := range characters {
		numbers[i], _ = strconv.Atoi(c)
	}
	return numbers
}

func max(numbers []int) (result int) {
	if len(numbers) > 0 {
		result = numbers[0]
	}
	for _, val := range numbers {
		if val > result {
			result = val
		}
	}
	return
}

func min(numbers []int) (result int) {
	if len(numbers) > 0 {
		result = numbers[0]
	}
	for _, val := range numbers {
		if val < result {
			result = val
		}
	}
	return
}

func rowChecksum(value string) int {
	numbers := stringToIntArray(value)
	diff := max(numbers) - min(numbers)
	return diff;
}

func Checksum(value string) int {
	rows := strings.Split(value, "\n")
	checksum := 0
	for _, row := range rows {
		checksum += rowChecksum(row)
	}
	return checksum
}
