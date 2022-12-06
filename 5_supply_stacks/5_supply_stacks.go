package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	// Note: for Windows line endings this needs to be \n\r\n
	// But Windows line endings are dumb so we switch input.txt to LF in vscode
	doot := strings.Split(text, "\n\n")

	fmt.Println(doot[0])
	fmt.Println("===")
	var crates = arrayifyCrateInput(doot[0])
	fmt.Println("OG:\n", crates)

	// Must copy to pass slice by value
	assumedCrates := make([]string, len(crates)*4)
	actualCrates := make([]string, len(crates)*4)
	copy(assumedCrates, crates)
	copy(actualCrates, crates)

	rearrangeCrates(assumedCrates, strings.Split(doot[1], "\n"), false)
	fmt.Println("Assumed:\n", assumedCrates)

	fmt.Println("OG:\n", crates)
	rearrangeCrates(actualCrates, strings.Split(doot[1], "\n"), true)
	fmt.Println("Actual:\n", actualCrates)

	assumed := ""
	actual := ""
	for i, _ := range crates {
		assumed += string(assumedCrates[i][len(assumedCrates[i])-1])
		actual += string(actualCrates[i][len(actualCrates[i])-1])
	}
	fmt.Printf("Assumed: %s\tActual: %s\n", assumed, actual)
}

func arrayifyCrateInput(input string) []string {
	var rows = strings.Split(input, "\n")
	var crates []string = make([]string, len(rows[0])/4+1)

	// This is slightly more efficient than just doing 2d array traversal
	// Since we only traverse up columns. It's negligible for the input range
	// and makes the code harder to read, but it's a good exercise for learning

	// For chars in bottom row (stack nums)
	for i := 0; i < len(rows[0]); i++ {
		c := string(rows[len(rows)-1][i])
		// fmt.Println(c)
		// If this is a crate column
		if c != " " && c != "\n" {
			// Move up the stack
			for j := (len(rows) - 2); j >= 0; j-- {
				// fmt.Printf("char at %d,%d is %s\n", j, i, string(rows[j][i]))
				if string(rows[j][i]) != " " {
					// Convert column index to array index
					cratei := i / 4
					crates[cratei] = string(crates[cratei]) + string(rows[j][i])
				} else {
					break
				}
			}
		}
	}
	fmt.Println(crates)
	return crates
}

func reverse(str string) (result string) {
	result = ""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func rearrangeCrates(crates []string, instructions []string, fancyCrane bool) []string {
	re := regexp.MustCompile("move\\s(\\d+).*(\\d+)\\sto\\s(\\d+)")
	for _, i := range instructions {
		// fmt.Println(i)
		parts := re.FindStringSubmatch(i)
		parts = []string{parts[1], parts[2], parts[3]}
		// fmt.Println(parts)
		if len(parts) > 0 {
			qty, _ := strconv.Atoi(parts[0])
			from, _ := strconv.Atoi(parts[1])
			to, _ := strconv.Atoi(parts[2])
			// Adj to arr index
			to -= 1
			from -= 1
			cratesToMove := crates[from][len(crates[from])-qty:]
			if fancyCrane {
				crates[to] += cratesToMove
			} else {
				crates[to] += reverse(cratesToMove)
			}
			crates[from] = string(crates[from][:len(crates[from])-qty])
			// fmt.Println(crates)
		}
	}
	return crates
}
