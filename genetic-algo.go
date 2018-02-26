package main

import (
	"fmt"
)

const TARGET_STR = "Hello World"
const MAX_ITERS = 10000
const CROSS_OVER_RATE = 0.7
const MUTATION_RATE = 0.45

func main() {
	population := makePopulation(10)
	fmt.Println("New population created: ")
	fmt.Println(population)

	var n_iters int = 0
	var children []Chromosome
	for n_iters < MAX_ITERS {
		var parents []Chromosome = pickBestM(population, 2)
		children = crossOver(parents[0], parents[1], CROSS_OVER_RATE)
		for i := range children {
			child := mutate(children[i], MUTATION_RATE)
			// 			fmt.Print("new child: ")
			// 			fmt.Println(child)
			addToPopulation(population, child)
		}
		n_iters += 1
	}

	fmt.Println("Current population")
	fmt.Println(population)
}
