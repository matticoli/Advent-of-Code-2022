package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stream := string(content)

	// Tune
	lastRepeat := 0

	// Check for next distinct string of length l - i, c are current index, char
	nextUnique := func(l int, i int, c byte) bool {
		// z is index lookback (for str of len 4, we check for dupes starting i - 3)
		z := l - 1
		if strings.Contains(stream[i-z:i], string(c)) {
			// Get char right after the last index of the repeated character
			rep := strings.LastIndex(stream[i-z:i], string(c)) + (i - z + 1)
			// Only update if there prev repetition isn't after this one
			if rep > lastRepeat {
				lastRepeat = rep
			}
		} else if i-lastRepeat >= z {
			fmt.Printf("Found sequence %s at %d\n", stream[i-z:i+1], i+1)
			lastRepeat = i
			return true
		}
		return false
	}

	tuned := false
	fmt.Println("Tuning...")
	for i := 3; i < len(stream); i++ {
		c := stream[i]
		if !tuned {
			if nextUnique(4, i, c) {
				tuned = true
				fmt.Println("Tuned. Awaiting message...")
			}
		} else {
			if nextUnique(14, i, c) {
				// Only the first message is important
				// We can just ignore the rest
				// It's probably fine
				break
			}
		}
	}
}
