package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Chromosome struct {
	gene    string
	fitness float64
}

/*
 "abc" => "011000010110001001100011"
*/
func stringToGene(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

/*
 "011000010110001001100011" => "abc"
*/
func geneToString(s string) string {
	var ans []string = make([]string, len(s)/8)
	for i := 0; i < len(s)/8; i++ {
		a, b := i*8, (i+1)*8
		tmp := s[a:b]
		x, err := strconv.ParseInt(tmp, 2, 64)
		if err == nil {
			ans = append(ans, string(x))
		}
	}
	return strings.Join(ans, "")
}

/*
 See how far the current string is from the target string
*/
func getFitness(cur string, target string) float64 {
	ans := float64(0)
	// for i := range target {
	// 	ans += math.Abs(float64(target[i]) - float64(cur[i]))
	// }
	// if ans == 0 {
	// 	return 1.0
	// } else {
	// 	return 1.0 / float64(ans)
	// }
	fmt.Printf("comparing: %s %s\n", cur, target)

	for i := range target {
		if target[i] != cur[i] {
			ans += 1
		}
	}
	return float64(len(target)) - ans
}

/*
 Flip one random bit along the gene as to mutate
*/
func mutate(chr Chromosome, rate float64) Chromosome {
	var index int = int(rate * float64(len(chr.gene)-1))
	var rep string = string(int('1') - int(chr.gene[index]) + int('0'))
	var ans string = chr.gene[0:index] + rep + chr.gene[index+1:len(chr.gene)]
	return Chromosome{ans, getFitness(ans, TARGET_STR)}
}

/*
 Create a new child from 2 parents
*/
func crossOver(pa1 Chromosome, pa2 Chromosome, rate float64) []Chromosome {
	var index int = int(rate * float64(len(pa1.gene)-1))
	var gene1 = pa1.gene[0:index] + pa2.gene[index:len(pa2.gene)]
	var gene2 = pa2.gene[0:index] + pa1.gene[index:len(pa1.gene)]
	var child1 Chromosome = Chromosome{gene1, getFitness(gene1, TARGET_STR)}
	var child2 Chromosome = Chromosome{gene2, getFitness(gene2, TARGET_STR)}
	return []Chromosome{child1, child2}
}
