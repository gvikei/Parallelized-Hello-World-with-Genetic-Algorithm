package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Population struct {
	genome []Chromosome
}

func makeChromosome(geneLen int, randomizer *rand.Rand) Chromosome {
	var ans []string = make([]string, geneLen)
	for i := 0; i < geneLen; i++ {
		ans[i] = string(randomizer.Int() % 128)
	}
	gene := strings.Join(ans, "")
	chr := Chromosome{gene, getFitness(gene, TARGET_STR)}
	fmt.Println(chr)
	return chr
}

func makePopulation(p Population, n int) {
	p.genome = make([]Chromosome, n)
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range p.genome {
		p.genome[i] = makeChromosome(len(TARGET_STR), randomizer)
	}
}

func pickBest(p Population, m int) {
	
	while m > 0 {
		m -= 1
		
	}
}