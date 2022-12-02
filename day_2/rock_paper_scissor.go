package day2

import (
	"bufio"
	"os"
)

func RockPaperScissorSumScores(rounds [][2]rune) int {
	sum := 0
	for _, round := range rounds {
		// round[1] is the shape you've selected
		ownSymbol := decipherSymbol(round[1])
		roundResult := ownSymbol + decipherOutcome(resolveOutcome(ownSymbol, decipherSymbol(round[0])))
		sum += roundResult
	}
	return sum
}

func RockPaperScissorSumDeterminedScores(rounds [][2]rune) int {
	sum := 0
	for _, round := range rounds {
		// round[1] is the shape you've selected
		ownSymbol := decipherChoice(decipherSymbol(round[0]), decipherSymbol(round[1]))
		roundResult := ownSymbol + decipherOutcome(resolveOutcome(ownSymbol, decipherSymbol(round[0])))
		sum += roundResult
	}
	return sum
}

func ParseStdInput() [][2]rune {
	scanner := bufio.NewScanner(os.Stdin)
	rounds := make([][2]rune, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "." {
			break
		}
		// str[0] and str[2] will be of interest
		rounds = append(rounds, [2]rune{rune(str[0]), rune(str[2])})
	}
	return rounds
}

// Inputs:
// - opponentChoice: either 1, 2, or 3 as in rock, paper, or scissor
// - expectedOutcome: either 1, 2, or 3 as in you need to lose, draw, or win
// returns your own choice this turn given the opponent's choice and such expected outcome
// returns 0 if there is something wrong
func decipherChoice(opponentChoice int, expectedOutcome int) int {
	switch opponentChoice {
	case 1:
		if expectedOutcome == 1 {
			// lose to rock
			return 3
		} else if expectedOutcome == 2 {
			// draw to rock
			return 1
		} else if expectedOutcome == 3 {
			// win to rock
			return 2
		}
	case 2:
		if expectedOutcome == 1 {
			// lose to paper
			return 1
		} else if expectedOutcome == 2 {
			// draw to paper
			return 2
		} else if expectedOutcome == 3 {
			// win to paper
			return 3
		}
	case 3:
		if expectedOutcome == 1 {
			// lose to scissor
			return 2
		} else if expectedOutcome == 2 {
			// draw to scissor
			return 3
		} else if expectedOutcome == 3 {
			// win to scissor
			return 1
		}
	default:
		return 0
	}
	return 0
}

// returns 1 for rock, 2 for paper, and 3 for scissor
// returns 1 for lose, 2 for draw, and 3 for win if the symbol is from the expected outcome
// returns 0 if unable to decipher the symbol
func decipherSymbol(symbol rune) int {
	switch symbol {
	case 'A', 'X':
		return 1
	case 'B', 'Y':
		return 2
	case 'C', 'Z':
		return 3
	default:
		return 0
	}
}

// Inputs:
// - outcome: -1 means lost, 1 means win, and 0 means draw
// returns 0 if unable to decipher the outcome
func decipherOutcome(outcome int) int {
	switch outcome {
	case -1:
		return 0
	case 1:
		return 6
	case 0:
		return 3
	default:
		return 0
	}
}

// Inputs:
// - a: your own move, either 1, 2, or 3
// - b: second player move
// returns -1, 1, or 0 depends on if the first player has lost, won, or draw against the second player
// returns 0 also for unidentified case
func resolveOutcome(a int, b int) int {
	switch {
	case a == b:
		return 0
	case a == 1:
		if b == 2 {
			// rock vs paper
			return -1
		} else if b == 3 {
			// rock vs scissor
			return 1
		}
	case a == 2:
		if b == 1 {
			// paper vs rock
			return 1
		} else if b == 3 {
			// paper vs scissor
			return -1
		}
	case a == 3:
		if b == 1 {
			// scissor vs rock
			return -1
		} else if b == 2 {
			// scissor vs paper
			return 1
		}
	default:
		return 0
	}
	return 0
}
