package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LanternFish int

var pool = []LanternFish{}

func (fish *LanternFish) Update() {
	if *fish == 0 {
		*fish = 6
		return
	}
	*fish -= 1
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	strNums, err := rd.ReadString('\n')
	strNums = strNums[:len(strNums)-1]
	if err != nil {
		panic(err)
	}
	tStrNums := strings.Split(strNums, ",")
	for _, strNum := range tStrNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		pool = append(pool, LanternFish(num))
	}
	for i := 0; i < 80; i++ {
		addFish := 0
		for fishIndex := 0; fishIndex < len(pool); fishIndex++ {
			fish := &pool[fishIndex]
			if *fish == 0 {
				addFish++
			}
			fish.Update()
		}
		for j := 0; j < addFish; j++ {
			pool = append(pool, 8)
		}
	}
	fmt.Println(len(pool))
}
