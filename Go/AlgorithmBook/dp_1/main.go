package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _, _ := reader.ReadLine()
		splittedText := strings.Split(string(text), " ")
		if len(splittedText) != 3 {
			break
		}
		m, _ := strconv.Atoi(splittedText[0])
		n, _ := strconv.Atoi(splittedText[1])
		t, _ := strconv.Atoi(splittedText[2])
		solve(m, n, t)
	}

}

// For recursive
func solveT(m, n, t int) int {
	if t == 0 {
		return 0
	}
	var first int
	if t >= m {
		first = solveT(m, n, t-m)
	} else {
		first = -1
	}

	var second int
	if t >= m {
		second = solveT(m, n, t-n)
	} else {
		second = -1
	}

	if first == -1 && second == -1 {
		return -1
	} else {
		return max(first, second) + 1
	}
}

func solve(m, n, t int) {
	var result, i int
	result = solveT(m, n, t)
	if result >= 0 {
		fmt.Printf("Zjedzono %d\n", result)
	} else {
		i = t - 1
		result = solveT(m, n, i)
		for result == -1 {
			i--
			result = solveT(m, n, i)
		}
		fmt.Printf("Zjedzono %d, pozostały czas %d\n", result, t-i)
	}
}
