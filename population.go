package main

import (
	"sort"
	"strings"
)

type byFitness []Chromosome

type Population struct {
	genome []Chromosome
}

func makeChromosome(geneLen int) Chromosome {
	var ans []string = make([]string, geneLen)
	for i := 0; i < geneLen; i++ {
		ans[i] = getRandomChar()
	}
	gene := strings.Join(ans, "")
	return Chromosome{gene, getFitness(gene, TARGET_STR)}
}

func makePopulation(n int) Population {
	var p = Population{make([]Chromosome, n)}
	// randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range p.genome {
		p.genome[i] = makeChromosome(len(TARGET_STR))
	}
	sort.Sort(byFitness(p.genome))
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
	return p.genome[0:m]
}

func addToPopulation(p Population, chr Chromosome) {
	p.genome[len(p.genome)-1] = chr
	sort.Sort(byFitness(p.genome))
}
