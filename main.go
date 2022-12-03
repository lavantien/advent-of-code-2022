package main

import (
	day1 "advent-of-code-2022/day_1"
	day2 "advent-of-code-2022/day_2"
	day3 "advent-of-code-2022/day_3"
	"fmt"
)

func main() {
	// firstDay()
	// secondDay()
	thirdDay()
}

func thirdDay() {
	items := day3.ParseStdInput()
	result := day3.RusksackFilter(items)
	fmt.Println(result)
}

func secondDay() {
	rounds := day2.ParseStdInput()
	result := day2.RockPaperScissorSumScores(rounds)
	fmt.Println(result)
	result2 := day2.RockPaperScissorSumDeterminedScores(rounds)
	fmt.Println(result2)
}

func firstDay() {
	pockets := day1.ParseStdInput()
	result := day1.CalorieCounting(pockets)
	fmt.Println(result)
	result2 := day1.TopKthSum(pockets, 3)
	fmt.Println(result2)
}
