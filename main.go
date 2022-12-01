package main

import (
	day1 "advent-of-code-2022/day_1"
	"fmt"
)

func firstDay() {
	pockets := day1.ParseStdInput()
	result := day1.CalorieCounting(pockets)
	fmt.Println(result)
	result2 := day1.TopKthSum(pockets, 3)
	fmt.Println(result2)
}

func main() {
	firstDay()
}
