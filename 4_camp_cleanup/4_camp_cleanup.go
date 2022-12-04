package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum, osum = countContained("./input.txt")
	// fmt.Printf("Contained: %d Overlap: %d\n", sum, osum)
}

func isContained(e1 string, e2 string) bool {
	var r1 = getRange(e1)
	var r2 = getRange(e2)
	// fmt.Printf("%s vs %s\n", e1, e2)
	if r1[0] <= r2[0] && r1[1] >= r2[1] {
		// R1 contains r2
		// fmt.Printf("r1 contains r2\n")
		return true
	} else if r1[0] >= r2[0] && r1[1] <= r2[1] {
		// R2 contains r1
		// fmt.Printf("r2 contains r1\n")
		return true
	} else {
		// fmt.Printf("not contained\n")
		return false
	}
}

func overlap(e1 string, e2 string) bool {
	var r1 = getRange(e1)
	var r2 = getRange(e2)
	// fmt.Printf("%s vs %s\n", e1, e2)
	if r1[0] >= r2[0] && r1[0] <= r2[1] {
		// R1 contains r2
		// fmt.Printf("r1 contains r2\n")
		return true
	} else if r2[0] >= r1[0] && r2[0] <= r1[1] {
		// R2 contains r1
		// fmt.Printf("r2 contains r1\n")
		return true
	} else {
		// fmt.Printf("no overlap\n")
		return false
	}
}

func getRange(e string) []int {
	var nums = strings.Split(e, "-")
	// Foolishly assume the input is valid and ignore errors 😇
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return []int{n1, n2}
}

func countContained(filename string) (int, int) {
	inputReader, error := os.Open(filename)
	if error != nil {
		fmt.Println("uh oh, no file")
		fmt.Println(error)
	}
	scanner := bufio.NewScanner(inputReader)

	var sum int = 0
	var osum int = 0
	for scanner.Scan() {
		var line = scanner.Text()
		var elves []string = strings.Split(line, ",")
		if isContained(elves[0], elves[1]) {
			sum++
		}
		if overlap(elves[0], elves[1]) {
			osum++
		}
	}
	inputReader.Close()
	return sum, osum
}
