package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	snowflakes := make([][]int, 0, 100000)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\r')
	c, _ := strconv.Atoi(strings.Replace(line, "\r", "", -1))

	for i := 0; i < c; i++ {
		line, _ := reader.ReadString('\r')
		curLine := strings.Replace(line, "\r", "", -1)
		snowflake := make([]int, 0, 6)
		for _, r := range curLine {
			if nb, err := strconv.Atoi(string(r)); err == nil {
				snowflake = append(snowflake, nb)
			}
		}
		snowflakes = append(snowflakes, snowflake)
	}

	identifyIdentical(snowflakes, len(snowflakes))
}

func identifyIdentical(a [][]int, n int) {
	var i, j int
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if identicalRight(a[i], a[j]) || identicalLeft(a[i], a[j]) {
				fmt.Println("Twin snowflakes found.")
				return
			}
		}
	}
	fmt.Println("No two snowflakes are alike.")
}

func identicalRight(first, second []int) bool {
	elIndex := -1
	for i, _ := range first {
		if first[i] == second[0] {
			elIndex = i
		}
	}
	if elIndex < 0 {
		return false
	}
	for i := 0; i < 6; i++ {
		if first[(elIndex+i)%6] != second[i] {
			return false
		}
	}
	return true
}

func identicalLeft(first, second []int) bool {
	elIndex := -1
	for i, _ := range first {
		if first[i] == second[0] {
			elIndex = i
		}
	}
	if elIndex < 0 {
		return false
	}
	for i := 0; i < 6; i++ {
		if first[(elIndex+i)%6] != second[(6-i)%6] {
			return false
		}
	}
	return true
}
