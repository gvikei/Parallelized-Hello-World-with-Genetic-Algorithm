package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var TARGET_STR = ""
var MAX_ITERS = 0
var INIT_N = 0
var XXX_TIMES = 0

var randomizer *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// var randomizer *rand.Rand = rand.New(rand.NewSource(100))

func _parse_args() {
	flag.StringVar(&TARGET_STR, "target", "Hello World", "Target string to be auto-generated")
	flag.IntVar(&MAX_ITERS, "max-iterations", 10000, "Max number of iterations to run the algorithm")
	flag.IntVar(&INIT_N, "population-size", 10, "The population size")
	flag.IntVar(&XXX_TIMES, "xxx-times", 1, "You know what this is")
	flag.Parse()
}

func main() {
	_parse_args()
	population := initPopulation(INIT_N)
	fmt.Println("Init population:")
	for i := range population.genome {
		fmt.Println(population.genome[i])
	}

	var n_iters int = 0
	var children []Chromosome
	var done = false
	for n_iters < MAX_ITERS && !done {
		var parents []Chromosome = pickBestM(population, 2)
		for j := 0; j < XXX_TIMES; j++ {
			children = crossOver(parents[0], parents[1], randomizer.Float64())
			for i := range children {
				child := mutate(children[i], randomizer.Float64())
				addToPopulation(population, child)
			}
		}
		n_iters += 1
		done = (strings.Compare(population.genome[0].gene, TARGET_STR) == 0)
	}

	fmt.Printf("Population after %d iterations: %s\n", n_iters, population.genome[0].gene)
	for i := range population.genome {
		fmt.Println(population.genome[i])
	}
}
