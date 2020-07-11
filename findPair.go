package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var sumTarget int = 9
	result := findCombination(numbers, sumTarget)
	fmt.Println(result)
}

func findCombination(numbers []int, target int) [][]int {
	result := [][]int{}
	if len(numbers) == 0 {
		return result
	}
	difference := target - numbers[0]
	if difference == 0 {
		result = append(result, []int{numbers[0]})
	} else if difference > 0 {
		result = findCombination(numbers[1:], difference)
		for index, value := range result {
			result[index] = append([]int{numbers[0]}, value...)
		}
	}
	for len(numbers) > 1 && numbers[0] == numbers[1] {
		numbers = numbers[1:]
	}
	result = append(result, findCombination(numbers[1:], target)...)
	return result
}
