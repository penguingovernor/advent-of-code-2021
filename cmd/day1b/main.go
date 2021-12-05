package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Sum(nums ...uint64) (result uint64) {
	for _, num := range nums {
		result += num
	}
	return
}

func main() {
	const windowSize = 3
	buffer := []uint64{}
	scanner := bufio.NewScanner(os.Stdin)
	increases := 0

	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, num)
		if len(buffer) < windowSize+1 {
			continue
		}
		window1 := Sum(buffer[:windowSize]...)
		window2 := Sum(buffer[1:]...)
		buffer = buffer[1:]
		if window2 > window1 {
			increases++
		}
	}

	fmt.Println(increases)
}
