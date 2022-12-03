package main

import (
	"bufio"
	"fmt"
	"os"
)

var scoring1 map[string]int32 = map[string]int32{
	// WLD + RPC
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

var scoring2 map[string]int32 = map[string]int32{
	// WLD + RPC
	"A X": 0 + 3,
	"A Y": 3 + 1,
	"A Z": 6 + 2,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 0 + 2,
	"C Y": 3 + 3,
	"C Z": 6 + 1,
}

func main() {
	var sum = totalScore("input.txt", scoring1)
	fmt.Printf("Pt 1 score is %d\n", sum)
	sum = totalScore("input.txt", scoring2)
	fmt.Printf("Pt 2 score is %d", sum)
}

func totalScore(filename string, scorematrix map[string]int32) int32 {
	inputReader, error := os.Open(filename)
	if error != nil {
		fmt.Println("uh oh, no file")
		fmt.Println(error)
	}
	scanner := bufio.NewScanner(inputReader)

	var sum int32 = 0
	for scanner.Scan() {
		var line = scanner.Text()
		sum += scorematrix[line]
	}
	inputReader.Close()
	return sum
}
