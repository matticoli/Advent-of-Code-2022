package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scorematrix string = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var sacks []string

func main() {
	var sum = totalScore("./input.txt")
	fmt.Printf("Pt 1 priority sum is %d\n", sum)
	fmt.Printf("Pt 2 priority sum is %d\n", badges())
}

func getPriority(item byte) int {
	return strings.Index(scorematrix, string(item))
}

func totalScore(filename string) int {
	inputReader, error := os.Open(filename)
	if error != nil {
		fmt.Println("uh oh, no file")
		fmt.Println(error)
	}
	scanner := bufio.NewScanner(inputReader)

	var sum int = 0
	for scanner.Scan() {
		var line = scanner.Text()
		sacks = append(sacks, line)
		sack1 := string(line[0 : len(line)/2])
		sack2 := string(line[len(line)/2 : len(line)])
		// fmt.Printf("line: %s\nsack1: %s\nsack2: %s\n", line, sack1, sack2)
		for i := 0; i < len(sack1); i++ {
			if strings.Contains(sack2, string(sack1[i])) {
				var priority = getPriority(sack1[i])
				// fmt.Printf("Found %s in %s\t Adding %d\n", string(sack1[i]), sack2, priority)
				sum += priority
				break
			}
		}
	}
	inputReader.Close()
	return sum
}

func badges() int {
	var sum int
	for i := 0; i < len(sacks); i += 3 {
		for j := 0; j < len(sacks[i]); j++ {
			if strings.Contains(sacks[i+1], string(sacks[i][j])) && strings.Contains(sacks[i+2], string(sacks[i][j])) {
				sum += getPriority(sacks[i][j])
				break
			}
		}
	}
	return sum
}
