package main

import (
	"math"
	"strings"
)

type Chromosome struct {
	gene    string
	fitness float64
}

/*
 See how far the current string is from the target string
*/
func getFitness(cur string, target string) float64 {
	ans := float64(0)
	for i := range target {
		if target[i] != cur[i] {
			ans += math.Abs(float64(target[i]) - float64(cur[i]))
		}
	}
	return -ans
}

func getRandomChar() string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	return string(chars[(randomizer.Int() % len(chars))])
}

/*
 Flip one random bit along the gene as to mutate
*/
func mutate(chr Chromosome, rate float64) Chromosome {
	var index int = len(chr.gene) - 1
	var tmp int = int(rate * float64(len(chr.gene)))
	if tmp < index {
		index = tmp
	}

	ok := false
	// Force mutation to be different from the previous gene
	var ans string
	for !ok {
		ans = chr.gene[0:index] + getRandomChar() + chr.gene[index+1:len(chr.gene)]
		ok = strings.Compare(ans, chr.gene) != 0
	}
	return Chromosome{ans, getFitness(ans, TARGET_STR)}
}

/*
 Create 2 chilren from 2 parents
*/
func crossOver(pa1 Chromosome, pa2 Chromosome, rate float64) []Chromosome {
	var index int = int(rate * float64(len(pa1.gene)-1))
	var gene1 = pa1.gene[0:index] + pa2.gene[index:len(pa2.gene)]
	var gene2 = pa2.gene[0:index] + pa1.gene[index:len(pa1.gene)]
	var child1 Chromosome = Chromosome{gene1, getFitness(gene1, TARGET_STR)}
	var child2 Chromosome = Chromosome{gene2, getFitness(gene2, TARGET_STR)}
	return []Chromosome{child1, child2}
}
