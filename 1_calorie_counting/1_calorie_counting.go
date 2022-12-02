package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var elfSums = fileToSums("input.txt")
	var cals = maxInt(elfSums)
	fmt.Printf("Top calorie elf has %d calories\n", cals)
	fmt.Printf("Top 3 elves have %d calories", sum(top3(elfSums)))
}

func maxInt(arr []int32) int32 {
	var cals int32 = 0
	for _, x := range arr {
		if x > cals {
			cals = x
		}
	}
	return cals
}

func top3(arr []int32) []int32 {
	var cals []int32 = []int32{0, 0, 0}
	for _, x := range arr {
		if x > cals[0] {
			cals[2] = cals[1]
			cals[1] = cals[0]
			cals[0] = x
		} else if x > cals[1] {
			cals[2] = cals[1]
			cals[1] = x
		} else if x > cals[2] {
			cals[2] = x
		}
	}
	return cals
}

func sum(arr []int32) int32 {
	var sum int32
	for _, x := range arr {
		sum += x
	}
	return sum
}

func fileToSums(filename string) []int32 {
	inputReader, error := os.Open(filename)
	if error != nil {
		fmt.Println("uh oh, no file")
		fmt.Println(error)
	}
	scanner := bufio.NewScanner(inputReader)

	var elfSums = []int32{}
	var sum int32 = 0
	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {
			num, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				fmt.Println(err)
			} else {
				sum += int32(num)
			}
		} else {
			elfSums = append(elfSums, sum)
			sum = 0
		}
	}
	inputReader.Close()
	return elfSums
}
