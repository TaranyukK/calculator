package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	userSelect       string
	userNumbers      string
	availableMethods = []string{"AVG", "SUM", "MED"}
)

func main() {
	fmt.Print("Выберите операцию (AVG, SUM, MED): ")
	getOperationSelect()
	fmt.Print("Введите числа через запятую: ")
	formattedNumbers := getNumbers()

	result := calculateNumbers(userSelect, formattedNumbers)

	fmt.Printf("Результат: %.2f", result)
}

func getOperationSelect() {
	for {
		fmt.Scanln(&userSelect)
		if slices.Contains(availableMethods, userSelect) {
			break
		} else {
			fmt.Print("Неизвестный метод (AVG, SUM, MED): ")
		}
	}
}

func getNumbers() []int64 {
	fmt.Scanln(&userNumbers)

	numbers := strings.Split(userNumbers, ",")
	result := make([]int64, len(numbers))
	for i, s := range numbers {
		v, _ := strconv.ParseInt(s, 10, 64)

		result[i] = v
	}

	return result
}

func calculateNumbers(method string, nums []int64) float64 {
	var result float64
	switch method {
	case "AVG":
		result = calculateAVG(nums)
	case "SUM":
		result = calculateSUM(nums)
	case "MED":
		result = calculateMED(nums)
	}

	return result
}

func calculateAVG(nums []int64) float64 {
	var sum float64

	for _, v := range nums {
		sum += float64(v)
	}

	return sum / float64(len(nums))
}

func calculateSUM(nums []int64) float64 {
	var result int64

	for _, s := range nums {
		result += s
	}

	return float64(result)
}

func calculateMED(nums []int64) float64 {
	slices.Sort(nums)
	mid := len(nums) / 2

	if len(nums)%2 == 0 {
		return (float64(nums[mid-1]) + float64(nums[mid])) / 2.0
	} else {
		return float64(nums[mid])
	}
}
