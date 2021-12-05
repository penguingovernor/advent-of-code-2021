package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BinaryCount struct {
	Zeros int
	Ones  int
}

func MergeCounts(a, b BinaryCount) BinaryCount {
	return BinaryCount{
		Zeros: a.Zeros + b.Zeros,
		Ones:  a.Ones + b.Ones,
	}
}

func GetMostFrequentInPosition(i int, m map[int]BinaryCount) rune {
	counts := m[i]
	if counts.Ones > counts.Zeros {
		return '1'
	}
	return '0'
}

func Opposite(c rune) rune {
	if c == '0' {
		return '1'
	}
	return '0'
}

func buildGammaRate(m map[int]BinaryCount) (uint64, error) {
	sb := strings.Builder{}
	for i := 0; i < len(m); i++ {
		sb.WriteRune(GetMostFrequentInPosition(i, m))
	}
	str := sb.String()
	return strconv.ParseUint(str, 2, 0)
}

func buildEpsilonRate(m map[int]BinaryCount) (uint64, error) {
	sb := strings.Builder{}
	for i := 0; i < len(m); i++ {
		sb.WriteRune(Opposite(GetMostFrequentInPosition(i, m)))
	}
	str := sb.String()
	return strconv.ParseUint(str, 2, 0)
}

func main() {
	counts := map[int]BinaryCount{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		binNum := scanner.Text()
		for i, bit := range binNum {
			current := BinaryCount{}
			if bit == '0' {
				current.Zeros++
			}
			if bit == '1' {
				current.Ones++
			}
			counts[i] = MergeCounts(counts[i], current)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	e, err := buildEpsilonRate(counts)
	if err != nil {
		panic(err)
	}
	g, err := buildGammaRate(counts)
	if err != nil {
		panic(err)
	}
	fmt.Println(g, e, g*e)
}
