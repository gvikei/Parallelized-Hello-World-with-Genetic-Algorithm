package main

import (
	"fmt"
)

const TARGET_STR = "Hello World"
const MAX_ITERS = 10
const CROSS_OVER_RATE = 0.7
const MUTATION_RATE = 0.45

func main() {
	population := makePopulation(10)
	for i := range population.genome {
		fmt.Println(population.genome[i])
	}

	var n_iters int = 0
	var children []Chromosome
	for n_iters < MAX_ITERS {
		var parents []Chromosome = pickBestM(population, 2)
		children = crossOver(parents[0], parents[1], CROSS_OVER_RATE)
		for i := range children {
			child := mutate(children[i], MUTATION_RATE)
			fmt.Printf("%s %d\n", child, len(child.gene))
			addToPopulation(population, child)
		}
		n_iters += 1
	}
}
