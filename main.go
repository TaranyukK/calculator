package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var availableMethods = map[string]func([]int64) float64{
	"AVG": calculateAVG,
	"SUM": calculateSUM,
	"MED": calculateMED,
}

func main() {
	fmt.Print("Выберите операцию (AVG, SUM, MED): ")
	selectedMethod := getOperationSelect()
	fmt.Print("Введите числа через запятую: ")
	formattedNumbers := getNumbers()

	result := selectedMethod(formattedNumbers)

	fmt.Printf("Результат: %.2f", result)
}

func getOperationSelect() func([]int64) float64 {
	var userSelect string
	for {
		fmt.Scanln(&userSelect)
		if _, ok := availableMethods[userSelect]; ok {
			break
		} else {
			fmt.Print("Неизвестный метод (AVG, SUM, MED): ")
		}
	}
	return availableMethods[userSelect]
}

func getNumbers() []int64 {
	var userNumbers string
	fmt.Scanln(&userNumbers)

	numbers := strings.Split(userNumbers, ",")
	result := make([]int64, len(numbers))
	for i, s := range numbers {
		v, _ := strconv.ParseInt(s, 10, 64)

		result[i] = v
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
