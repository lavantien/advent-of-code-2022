package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Arguments:
// - pockets: pockets[i] is the pocket of the i-th Elf, which contains a list of Calorie numbers
// Returns:
// - total Calories of the Elf that carrying the most Calories
func CalorieCounting(pockets [][]int) int {
	maxSum := math.MinInt
	for _, pocket := range pockets {
		localSum := 0
		for _, calorie := range pocket {
			localSum += calorie
		}
		if localSum > maxSum {
			maxSum = localSum
		}
	}
	return maxSum
}

// Returns:
// - total Calories of the top k Elves that carryign the most Calories
func TopKthSum(pockets [][]int, k int) int {
	// We can just reflect each pocket to its sum of Calories and then sort and then get sum of the top k
	sums := []int{}
	for _, pocket := range pockets {
		localSum := 0
		for _, calorie := range pocket {
			localSum += calorie
		}
		sums = append(sums, localSum)
	}
	sort.SliceStable(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	sum := 0
	if k > len(sums) {
		k = len(sums)
	}
	fmt.Println(sums) // debug
	for i := 0; i < k; i++ {
		sum += sums[i]
	}
	return sum
}

// Let the final "\n." be the exit signal
func ParseStdInput() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	pockets := make([][]int, 0)
	calories := []int{}
	for scanner.Scan() {
		str := scanner.Text()
		if str == "." {
			break
		}
		if str == "" {
			pockets = append(pockets, calories)
			calories = nil
		} else {
			var calorie int
			fmt.Sscan(str, &calorie)
			calories = append(calories, calorie)
		}
	}
	return pockets
}
