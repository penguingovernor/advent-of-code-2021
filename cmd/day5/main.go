package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GCDRemainder calculates GCD iteratively using remainder.
func GCDRemainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

type Point struct {
	X, Y int
}

func (p Point) Key() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func Equal(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func PointFromString(s string) Point {
	fields := strings.Split(s, ",")
	x, err := strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	return Point{X: x, Y: y}
}

func EnumerateIntermediaryPoints(p1, p2 Point) <-chan Point {
	// Always go from p1 to p2
	var update func(p *Point)

	// L->R
	if p1.Y == p2.Y && p1.X < p2.X {
		update = func(p *Point) { p.X += 1 }
	}
	// R->L
	if p1.Y == p2.Y && p1.X > p2.X {
		update = func(p *Point) { p.X -= 1 }
	}
	// D->U
	if p1.X == p2.X && p1.Y < p2.Y {
		update = func(p *Point) { p.Y += 1 }
	}
	// U->D
	if p1.X == p2.X && p1.Y > p2.Y {
		update = func(p *Point) { p.Y -= 1 }
	}

	points := make(chan Point)

	go func() {
		runner := p1
		for !Equal(runner, p2) {
			points <- runner
			update(&runner)
		}
		points <- runner
		close(points)
	}()

	return points
}

func main() {
	pointMap := map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), "->")
		points[0] = strings.TrimSpace(points[0])
		points[1] = strings.TrimSpace(points[1])
		p1 := PointFromString(points[0])
		p2 := PointFromString(points[1])

		if !(p1.X == p2.X || p1.Y == p2.Y) {
			continue
		}

		for point := range EnumerateIntermediaryPoints(p1, p2) {
			pointMap[point.Key()]++
		}
	}

	var result int
	for _, v := range pointMap {
		if v >= 2 {
			result++
		}
	}

	fmt.Println(result)
}
