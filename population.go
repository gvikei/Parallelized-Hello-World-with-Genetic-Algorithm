package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type byFitness []Chromosome

type Population struct {
	genome []Chromosome
}

func makeChromosome(geneLen int, randomizer *rand.Rand) Chromosome {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	var ans []rune = make([]rune, geneLen)
	for i := 0; i < geneLen; i++ {
		ans[i] = chars[(randomizer.Int() % len(chars))]
	}
	gene := string(ans)
	chr := Chromosome{gene, getFitness(gene, TARGET_STR)}
	fmt.Println(chr)
	return chr
}

func makePopulation(n int) Population {
	var p = Population{make([]Chromosome, n)}
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range p.genome {
		p.genome[i] = makeChromosome(len(TARGET_STR), randomizer)
	}
	return p
}

func (s byFitness) Len() int {
	return len(s)
}
func (s byFitness) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byFitness) Less(i, j int) bool {
	// Reverse order - chromosome with larger fitness comes first
	return s[i].fitness > s[j].fitness
}

func pickBestM(p Population, m int) []Chromosome {
	sort.Sort(byFitness(p.genome))
	fmt.Println(p.genome)
	return p.genome[0:m]
}

func addToPopulation(p Population, chr Chromosome) {
	p.genome[len(p.genome)-1] = chr
	sort.Sort(byFitness(p.genome))
}
