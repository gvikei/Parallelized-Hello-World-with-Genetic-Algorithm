package main

import (
    "fmt"
)

type Chromosome struct {
    gene string
}

func mutate(chr Chromosome, rate float64) Chromosome {
    var index int = int( rate * float64(len(chr.gene) - 1) )
    var rep string = string(int('1') - int(chr.gene[index]) + int('0'))
    var ans string =
            chr.gene[0:index] + rep + chr.gene[index+1:len(chr.gene)]
    return Chromosome{ans}
}

func crossOver(pa1 Chromosome, pa2 Chromosome, rate float64) []Chromosome {
    var index int = int( rate * float64(len(pa1.gene) - 1) )
    
    var child1 Chromosome = 
        Chromosome{pa1.gene[0:index] + pa2.gene[index+1:len(pa2.gene)]}
    var child2 Chromosome = 
        Chromosome{pa2.gene[0:index] + pa1.gene[index+1:len(pa1.gene)]}
        
    // fmt.Printf(
    //     "%s & %s ==> %s, %s\n", pa1.gene, pa2.gene, child1.gene, child2.gene)
        
    return []Chromosome{child1, child2}
}

func main() {
    var c1 Chromosome = Chromosome{"101"}
    mutate(c1, 0.7)
    // var c2 Chromosome = Chromosome{"1011010001"}
    // fmt.Println(crossOver(c1, c2, 0.7))
}