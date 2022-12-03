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

func RusksackFilterTriplet(items [][3]string) int {
	sum := 0
	for _, item := range items {
		table1 := make(map[rune]int)
		table2 := make(map[rune]int)
		table3 := make(map[rune]int)
		runes1 := []rune(item[0])
		runes2 := []rune(item[1])
		runes3 := []rune(item[2])
		for _, rune := range runes1 {
			table1[rune]++
		}
		for _, rune := range runes2 {
			table2[rune]++
		}
		for _, rune := range runes3 {
			table3[rune]++
		}
		for k1 := range table1 {
			for k2 := range table2 {
				for k3 := range table3 {
					if k1 == k2 && k1 == k3 && k2 == k3 {
						priority := decodePriority(k1)
						fmt.Println(k1, string(k1), priority) // debug
						sum += priority
						break
					}
				}
			}
		}
	}
	return sum
}

func ParseStdInput2(input1 []string) [][3]string {
	items := make([][3]string, 0)
	counter := 0
	triplet := [3]string{}
	for _, str := range input1 {
		triplet[counter] = str
		counter++
		if counter == 3 {
			counter = 0
			items = append(items, triplet)
		}
	}
	return items
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
