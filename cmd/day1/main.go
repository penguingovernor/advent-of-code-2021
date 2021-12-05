package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	prevValue := uint64(math.MaxUint64)
	increases := 0

	// Create a scanner.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Get the number.
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		// Is this number smaller than or equal the prev. value?
		if num > prevValue {
			increases++
		}

		prevValue = num
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(increases)
}
