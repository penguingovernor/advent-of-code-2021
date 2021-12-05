package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var depth = 0
var horizontal = 0
var aim = 0

func Forward(x int) {
	horizontal += x
	depth += (aim * x)
}

func Down(x int) {
	aim += x
}

func Up(x int) {
	aim -= x
}

var commandToFuncMac = map[string]func(int){
	"forward": Forward,
	"down":    Down,
	"up":      Up,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			panic(fmt.Sprintf("wrong number of inputs for %q", scanner.Text()))
		}
		cmd := fields[0]
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		cmdFunc, ok := commandToFuncMac[cmd]
		if !ok {
			panic(fmt.Sprintf("unknown command %q", cmd))
		}
		cmdFunc(n)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(horizontal * depth)
}
