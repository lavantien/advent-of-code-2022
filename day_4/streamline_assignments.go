package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Inputs:
// - assignments: [0]-[1],[2]-[3] reflect a-b,c-d where a to b is the sections assigned to the first elf in pair
// returns the number of assignment pairs where one range fully contains the other
func StreamlineAssignments(assignments [][4]int) int {
	count := 0
	for _, assignment := range assignments {
		if mutualContain(assignment) {
			fmt.Println(assignment) // debug
			count++
		}
	}
	return count
}

// returns true if, for any pair a-b,c-d, a-b contains c-d or vice versa
func mutualContain(assignment [4]int) bool {
	firstContainsSecond := assignment[0] <= assignment[2] && assignment[1] >= assignment[3]
	secondContainsFirst := assignment[2] <= assignment[0] && assignment[3] >= assignment[1]
	return firstContainsSecond || secondContainsFirst
}

func ParseStdInput() [][4]int {
	scanner := bufio.NewScanner(os.Stdin)
	assignments := make([][4]int, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "." {
			break
		}
		// line input example: 2-8,3-7
		tokens := strings.FieldsFunc(str, func(r rune) bool {
			return r == '-' || r == ','
		})
		var assignment [4]int
		for i, token := range tokens {
			fmt.Sscan(token, &assignment[i])
		}
		assignments = append(assignments, assignment)
	}
	return assignments
}
