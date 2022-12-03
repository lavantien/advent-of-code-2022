package day3

import (
	"bufio"
	"fmt"
	"os"
)

// Inputs:
// - items: composed of a list of rusksacks that contains lower case and upper case characters
// returns the sum of share item's priorities across all rusksack
func RusksackFilter(items []string) int {
	sum := 0
	for _, item := range items {
		table1 := make(map[rune]int)
		table2 := make(map[rune]int)
		runes := []rune(item)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			table1[runes[i]]++
			table2[runes[j]]++
		}
		for k1 := range table1 {
			for k2 := range table2 {
				if k1 == k2 {
					priority := decodePriority(k1)
					fmt.Println(k1, string(k1), priority) // debug
					sum += priority
					break
				}
			}
		}
	}
	return sum
}

func ParseStdInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	items := []string{}
	for scanner.Scan() {
		str := scanner.Text()
		if str == "." {
			break
		}
		items = append(items, str)
	}
	return items
}

// returns 0 if item is broken
func decodePriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 96)
	} else if item >= 'A' && item <= 'Z' {
		return int(item - 64 + 26)
	}
	return 0
}
