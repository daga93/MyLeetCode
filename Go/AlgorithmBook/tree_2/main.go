package main

import "fmt"

type line struct {
	start int
	stop  int
}

type Climber struct {
	currentHeight int
	jumpHeight    int
	currentMoves  int
}

func main() {
	// Climb to h meters
	// Climbing starts at 0m
	// At start, he can jump to j meters, that is the only jump.
	// Powder Effects - he can't jump j, he can't go down if he went to another Poweder

	h, j, _ := 15, 4, 3 // h - height to reach, j - single jump height, l - number of segments with powder
	powderedElements := []line{
		{start: 2, stop: 3},
		{start: 7, stop: 8},
		{start: 13, stop: 14},
	}

	possiblePositions := []int{0}
	move := 0

climbing:
	for {
		possiblePositions = makeMove(powderedElements, possiblePositions, j)
		move++

		fmt.Println("After move:", move)
		fmt.Println("Possible positions:", possiblePositions)

		for _, posHeight := range possiblePositions {
			if posHeight >= h {
				fmt.Println("Jumped to heigh in moves:", move)
				break climbing
			}
		}
	}

}

func makeMove(powderLines []line, currentPos []int, jumpHeight int) []int {
	jumpMoves := jump(powderLines, currentPos, jumpHeight)
	goDownMoves := []int{}
	for _, pos := range currentPos {
		goDownMoves = append(goDownMoves, goDown(powderLines, pos, jumpHeight)...)
	}
	return append(jumpMoves, goDownMoves...)
}

func jump(powderLines []line, currentPos []int, jumpHeight int) []int {
	possPos := []int{}
	for _, pos := range currentPos {
		if pos == 0 {
			possPos = append(possPos, jumpHeight)
		} else {
			if !isInPowder(powderLines, pos) {
				possPos = append(possPos, pos+jumpHeight)
			}
		}
	}
	return possPos
}

func goDown(powderLines []line, currentPos int, jumpHeight int) []int {
	possiblePositions := []int{}
	if isInPowder(powderLines, currentPos) {
		for i := 1; i < jumpHeight; i++ {
			if !isInPowder(powderLines, currentPos-i) && currentPos-i > 0 {
				possiblePositions = append(possiblePositions, currentPos-i)
			}
		}
	} else {
		for i := 0; i < jumpHeight; i++ {
			if !isInPowder(powderLines, currentPos-i) && currentPos-i > 0 {
				possiblePositions = append(possiblePositions, currentPos-i)
			}
		}
	}
	return possiblePositions
}

func isInPowder(powderedLines []line, height int) bool {
	for _, l := range powderedLines {
		if height >= l.start && height <= l.stop {
			return true
		}
	}
	return false
}
