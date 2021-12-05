package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Filter(binStrings []string, position int, compare func(zeroCount, oneCount int) rune) []string {
	zeroCount := 0
	oneCount := 0
	for _, binStr := range binStrings {
		switch {
		case binStr[position] == '0':
			zeroCount++
		case binStr[position] == '1':
			oneCount++
		}
	}
	winner := compare(zeroCount, oneCount)
	result := []string{}
	for _, binStr := range binStrings {
		if binStr[position] == byte(winner) {
			result = append(result, binStr)
		}
	}

	return result
}

func BinStrings(rd io.Reader) ([]string, error) {
	results := []string{}
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}
	return results, scanner.Err()
}

func Oxygen(zeroCount, oneCount int) rune {
	switch {
	case zeroCount > oneCount:
		return '0'
	case oneCount > zeroCount:
		return '1'
	default:
		return '1'
	}
}

func Co2(zeroCount, oneCount int) rune {
	switch {
	case zeroCount < oneCount:
		return '0'
	case oneCount < zeroCount:
		return '1'
	default:
		return '0'
	}
}

func Code(binStr []string, compare func(zeroCount, oneCount int) rune) string {
	winner := ""
	pos := 0
	for {
		newStrs := Filter(binStr, pos, compare)
		if len(newStrs) == 1 {
			winner = newStrs[0]
			break
		}
		pos++
		binStr = newStrs
	}
	return winner
}

func Answer(oxygen, co2 string) uint64 {
	o, err := strconv.ParseUint(oxygen, 2, 0)
	if err != nil {
		panic(err)
	}
	c, err := strconv.ParseUint(co2, 2, 0)
	if err != nil {
		panic(err)
	}
	return o * c
}

func main() {
	binStrings, err := BinStrings(os.Stdin)
	if err != nil {
		panic(err)
	}
	oxygen := Code(binStrings, Oxygen)
	co2 := Code(binStrings, Co2)

	fmt.Println(oxygen, co2, Answer(oxygen, co2))
}
