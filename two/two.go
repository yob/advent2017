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

func filterDivisible(numbers []int) (int, int) {
	for _, outerVal := range numbers {
		for _, innerVal := range numbers {
			if innerVal != outerVal && innerVal % outerVal == 0 {
				return outerVal, innerVal
			}
		}
	}
	return 0, 0
}

func rowChecksum(value string) int {
	numbers := stringToIntArray(value)
	diff := max(numbers) - min(numbers)
	return diff;
}

func rowDivisibleChecksum(value string) int {
	numbers := stringToIntArray(value)
	low, high := filterDivisible(numbers)
	if high == 0 || low == 0 {
		return 0
	} else {
		return high / low;
	}
}

func Checksum(value string) int {
	rows := strings.Split(value, "\n")
	checksum := 0
	for _, row := range rows {
		checksum += rowChecksum(row)
	}
	return checksum
}

func ChecksumDivisible(value string) int {
	rows := strings.Split(value, "\n")
	checksum := 0
	for _, row := range rows {
		checksum += rowDivisibleChecksum(row)
	}
	return checksum
}
