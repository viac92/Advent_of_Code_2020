// light red bags contain 1 bright white bag, 2 muted yellow bags.
// dark orange bags contain 3 bright white bags, 4 muted yellow bags.
// bright white bags contain 1 shiny gold bag.
// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
// dark olive bags contain 3 faded blue bags, 4 dotted black bags.
// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// faded blue bags contain no other bags.
// dotted black bags contain no other bags.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
)

type bag struct {
	numberOfBag int
	bagName string
}

type bagRule struct {
	ruleName string
	bagConteined []bag
}

func parseLine(line string) (bRule bagRule, err error) {
	var b bag
	var nameSlice []string
	nameBagRule := regexp.MustCompile("([a-z])+")
	nameSlice = nameBagRule.FindAllString(line, 2)
	bRule.ruleName = nameSlice[0] + " " + nameSlice[1]


	if strings.ContainsAny(line, "0123456789") {
		
		numberBagConteinde := regexp.MustCompile("[0-9]")	
		nameBagContained := regexp.MustCompile("([a-z])+")
		for strings.ContainsAny(line, "0123456789") {
			
			b.numberOfBag,_ = strconv.Atoi(numberBagConteinde.FindString(line))			
			line = line[numberBagConteinde.FindStringIndex(line)[0] + 1:]
			s := nameBagContained.FindAllString(line, 2) 
			b.bagName = s[0] + " " + s[1] 
					
			bRule.bagConteined = append(bRule.bagConteined, b)
		}
	}
	return
}

func main() {
	b := bufio.NewScanner(os.Stdin)
	var bRule []bagRule
	var okBag []string
	var okBagCounter int

	for b.Scan() {
		t := b.Text()
		bR,_ := parseLine(t) 
		bRule = append(bRule, bR)
	}

	for _,bR := range bRule {
		for _,bg := range bR.bagConteined {
			if bg.bagName == "shiny gold" {
				okBag = append(okBag, bR.ruleName)
				okBagCounter++
				break
			}
		}
	}

	var count int
	for {
		fmt.Println(okBag[count:])
		for _,ok := range okBag[count:] {
			count++
			for _,bR := range bRule {
				loop: for _,bg := range bR.bagConteined {
					if ok == bg.bagName {
						for _,allOk := range okBag {
							if allOk == bR.ruleName {
								break loop
							}
						}
						okBag = append(okBag, bR.ruleName)
						okBagCounter++
						break
					}
				}
			}
		}
		if count == len(okBag) {
			break
		}
		fmt.Println(count)
	}

	// fmt.Println(bRule)
	// fmt.Println()
	for _,r := range okBag {
	 fmt.Println(r)
	}
	fmt.Println(len(okBag))
	// fmt.Println(okBagCounter)
}
