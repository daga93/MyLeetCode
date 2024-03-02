package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

type Board [][]int

type Positions []Position

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	splittedText := strings.Split(string(text), " ")
	if len(splittedText) != 6 {
		return
	}
	planeWidth, _ := strconv.Atoi(splittedText[0])
	planeHeight, _ := strconv.Atoi(splittedText[1])
	pRow, _ := strconv.Atoi(splittedText[2])
	pCol, _ := strconv.Atoi(splittedText[3])
	kRow, _ := strconv.Atoi(splittedText[4])
	kCol, _ := strconv.Atoi(splittedText[5])
	fmt.Println("Data", planeWidth, planeHeight, pRow, pCol, kRow, kCol)
}

func findDistance(kRow, kCol, destRow, destCol, numRows, numCols int) int {
	var curPositions, newPositions Positions
	var numCurPositions, numNewPositions int
	var minMoves Board
	for i := 1; i <= numRows; i++ {
		for j := 1; j <= numCols; j++ {
			minMoves[i][j] = -1
		}
	}
	minMoves[kRow][kCol] = 0
	curPositions = append(curPositions, Position{row: kRow, col: kCol})
	numCurPositions = 1
	for numCurPositions > 0 {
		numNewPositions = 0
		for i := 0; i < numCurPositions; i++ {
			fromRow := curPositions[i].row
			fromCol := curPositions[i].col
			if fromRow == destRow && fromCol == destCol {
				return minMoves[destRow][destCol]
			}
			addPosition(fromRow, fromCol, fromRow+1, fromCol+2, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow+1, fromCol-2, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow-1, fromCol+2, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow-1, fromCol-2, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow+2, fromCol+1, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow+2, fromCol-1, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow-2, fromCol+1, numRows, numCols, newPositions, &numNewPositions, minMoves)
			addPosition(fromRow, fromCol, fromRow-2, fromCol-1, numRows, numCols, newPositions, &numNewPositions, minMoves)
		}
		numCurPositions = numNewPositions
		for _, curPos := range newPositions {
			curPositions = append(curPositions, curPos)
		}
	}
	return -1
}

func addPosition(fromRow, fromCol, toRow, toCol, numRows, numCols int, newPositions Positions, numNewPositions *int, minMoves Board) {
	if toRow >= 1 && toCol >= 1 && toRow <= numRows && toCol <= numCols && minMoves[toRow][toCol] == -1 {
		minMoves[toRow][toCol] = 1 + minMoves[fromRow][fromCol]
		newPosition := Position{row: toRow, col: toCol}
		newPositions = append(newPositions, newPosition)
		*numNewPositions++
	}
}
