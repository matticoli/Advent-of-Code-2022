package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func num(i byte) int {
	res, _ := strconv.ParseInt(string(i), 10, 0)
	return int(res)
}

func printmap(arr [][]bool) {
	count := 0
	for i, _ := range arr {
		for j, _ := range arr[i] {
			if arr[i][j] {
				fmt.Print("x")
				count++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stream := string(content)
	// Map of tree heights:
	grid := strings.Split(stream, "\n")
	// Working map of visible trees:
	b := make([][]bool, len(grid))

	// Func to get index from end of arr
	var rev = func(index int) int {
		return len(grid) - index - 1
	}

	var g = func(i int, j int) int {
		return num(grid[i][j])
	}

	for i := range grid {
		tr := g(i, 0)
		tl := g(i, rev(0))
		td := g(0, i)
		tu := g(rev(0), i)
		// yay malloc
		if b[i] == nil {
			b[i] = make([]bool, len(grid))
		}
		if b[rev(i)] == nil {
			b[rev(i)] = make([]bool, len(grid))
		}
		b[i][0] = true
		b[i][rev(0)] = true
		b[0][i] = true
		b[rev(0)][i] = true

		for j := range grid {
			if b[j] == nil {
				b[j] = make([]bool, len(grid))
			}
			if g(i, j) > tr {
				tr = g(i, j)
				b[i][j] = true
			}
			if g(i, rev(j)) > tl {
				tl = g(i, rev(j))
				b[i][rev(j)] = true
			}
			if g(j, i) > td {
				td = g(j, i)
				b[j][i] = true
			}
			if g(j, rev(i)) > tu {
				tu = g(j, rev(i))
				b[j][rev(i)] = true
			}
			// Suppress unused (why tho??):
			tr = tr
			tl = tl
			tu = tu
			td = td
		}
	}
	printmap(b)
}
