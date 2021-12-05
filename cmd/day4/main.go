package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type BingoCell struct {
	Value string
	Seen  bool
}

type BingoBoard struct {
	Board  [][]BingoCell
	Lookup map[string][2]int
}

func NewBingoBoard(lines [5]string) BingoBoard {
	b := BingoBoard{}
	b.Lookup = make(map[string][2]int)
	row := 0

	for _, bingoRow := range lines {
		values := strings.Fields(bingoRow)
		cells := make([]BingoCell, 5)
		col := 0
		for i := 0; i < len(values); i++ {
			cells[i] = BingoCell{Value: values[i], Seen: false}
			b.Lookup[values[i]] = [2]int{row, col}
			col++
		}
		row++
		b.Board = append(b.Board, cells)
	}

	return b
}

func (b *BingoBoard) Mark(value string) {
	rowCol, ok := b.Lookup[value]
	if !ok {
		return
	}
	b.Board[rowCol[0]][rowCol[1]].Seen = true
}

func (b *BingoBoard) Winner() bool {
	rowIsWinner := func(r []BingoCell) bool {
		for _, c := range r {
			if !c.Seen {
				return false
			}
		}
		return true
	}
	colIsWinner := func(c0, c1, c2, c3, c4 BingoCell) bool {
		return (c0.Seen && c1.Seen && c2.Seen && c3.Seen && c4.Seen)
	}

	for i := 0; i < 5; i++ {
		if rowIsWinner(b.Board[i]) {
			return true
		}
		colWinner := colIsWinner(
			b.Board[0][i],
			b.Board[1][i],
			b.Board[2][i],
			b.Board[3][i],
			b.Board[4][i],
		)
		if colWinner {
			return true
		}
	}

	return false
}

func (b *BingoBoard) Value() (result int) {
	for i := 0; i < len(b.Board); i++ {
		for j := 0; j < len(b.Board[i]); j++ {
			if !b.Board[i][j].Seen {
				n, err := strconv.Atoi(b.Board[i][j].Value)
				if err != nil {
					panic(err)
				}
				result += n
			}
		}
	}
	return result
}

func handleScannerError(scanner *bufio.Scanner) {
	err := scanner.Err()
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	handleScannerError(scanner)
	numbers := strings.Split(scanner.Text(), ",")

	doEmpty := true
	count := 0
	lines := [5]string{}
	boards := []BingoBoard{}
	for scanner.Scan() {
		if doEmpty {
			doEmpty = false
			continue
		}
		lines[count] = scanner.Text()
		count++

		if count == 5 {
			count = 0
			boards = append(boards, NewBingoBoard(lines))
			doEmpty = true
		}
	}
	handleScannerError(scanner)

	winningNum := ""
	winningValue := 0

outer:
	for _, num := range numbers {
		for _, b := range boards {
			b.Mark(num)
			if b.Winner() {
				winningValue = b.Value()
				winningNum = num
				break outer
			}
		}
	}

	n, err := strconv.Atoi(winningNum)
	if err != nil {
		panic(err)
	}

	fmt.Println(winningValue * n)
}
