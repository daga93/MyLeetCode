package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	snowflakes := map[int][][]int{}

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
		key := hashKey(snowflake)
		if _, ok := snowflakes[key]; !ok {
			snowflakes[key] = [][]int{snowflake}
		} else {
			snowflakes[key] = append(snowflakes[key], snowflake)
		}
	}
	fmt.Println(snowflakes)
	found := false
	for _, itms := range snowflakes {
		if len(itms) > 1 {
			found = identifyIdentical(itms, len(itms))
		}
	}
	if found {
		fmt.Println("Twin snowflakes found.")
	} else {
		fmt.Println("No two snowflakes are alike.")
	}
}

func hashKey(snowflake []int) int {
	return (snowflake[0] + snowflake[1] + snowflake[2] + snowflake[3] + snowflake[4] + snowflake[5]) % 100000
}

func identifyIdentical(a [][]int, n int) bool {
	var i, j int
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if identicalRight(a[i], a[j]) || identicalLeft(a[i], a[j]) {
				return true
			}
		}
	}
	return false
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
