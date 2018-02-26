package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const TARGET_STR = "Hello World"
const MAX_ITERS = 100000
const CROSS_OVER_RATE = 0.7
const MUTATION_RATE = 0.45

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	population := makePopulation(10)
	fmt.Println("New population created: ")
	fmt.Println(population)

	var n_iters int = 0
	var children []Chromosome
	var done = false
	for n_iters < MAX_ITERS || done {
		var parents []Chromosome = pickBestM(population, 2)
		children = crossOver(parents[0], parents[1], randomizer.Float64())
		for i := range children {
			child := mutate(children[i], randomizer.Float64())
			addToPopulation(population, child)
		}
		n_iters += 1
		done = (strings.Compare(population.genome[0].gene, TARGET_STR) == 0)
	}

	fmt.Printf("Current population after %d iterations:\n", n_iters)
	fmt.Println(population)
}
