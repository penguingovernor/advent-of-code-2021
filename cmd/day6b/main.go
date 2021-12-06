package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Can't keep track of all fish in huge pool
// Even with bit optimizations (2 fish per byte, where any non 0d0-0d8 are considered invalid) wouldn't be sufficient.
// Only keep track of how much of each type of fish/timer we actually have.
type FishPool [9]*big.Int

func NewFishPool() FishPool {
	fp := FishPool{}
	for i := 0; i < len(fp); i++ {
		fp[i] = new(big.Int).SetInt64(0)
	}
	return fp
}

func (pool *FishPool) UpdatePool() {
	// The number of zero timer fish we have is the number of new fish we're adding.
	// New fish are represetned at index 8.
	newFish := new(big.Int).Set(pool[0])

	// Then we shift down our counts, ie slot 0 get's slot's 1 contents, slot 1 slot's 2, etc
	// All the way to 7 <- 8
	for i := 0; i < len(pool)-1; i++ {
		pool[i].Set(pool[i+1])
	}

	// Fish's with count 0 are reset to 6
	pool[6].Add(pool[6], newFish)
	// Set the number of new to the number of 0 fish
	pool[8].Set(newFish)
}

func (pool *FishPool) IncrementPosition(n int) {
	one := new(big.Int).SetInt64(1)
	pool[n].Add(pool[n], one)
}

func (pool *FishPool) Sum() *big.Int {
	total := new(big.Int).SetInt64(0)
	for _, fishes := range pool {
		total.Add(total, fishes)
	}
	return total
}

func main() {

	// Get the input.
	rd := bufio.NewReader(os.Stdin)
	strNums, err := rd.ReadString('\n')
	strNums = strNums[:len(strNums)-1]
	if err != nil {
		panic(err)
	}
	tStrNums := strings.Split(strNums, ",")

	// Populate the pool.
	pool := NewFishPool()
	for _, strNum := range tStrNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		pool.IncrementPosition(num)
	}

	// Count the fish.
	for i := 0; i < 256; i++ {
		pool.UpdatePool()
	}

	// Print the fish
	fmt.Println(pool.Sum())
}
