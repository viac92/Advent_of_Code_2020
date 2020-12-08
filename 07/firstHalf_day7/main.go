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
	"strconv"
	"strings"
)

type bag struct {
	num  int
	name string
}

func createBag(str []string) []bag {
	var bagOut []bag
	var num int
	if len(str) == 7 {
		bagOut = append(bagOut, bag{0, str[1] + " " + str[0]})
		return bagOut
	}
	for i := len(str) - 2; i > 0; i -= 4 {
		if i == 1 {
			num = 0
		} else {
			num, _ = strconv.Atoi(str[i-2])
		}
		bagOut = append(bagOut, bag{num, str[i-1] + " " + str[i]})
		if i == 6 {
			i--
		}
	}
	return bagOut
}

func main() {
	b := bufio.NewScanner(os.Stdin)
	var bagSlice []bag
	var canGold []bag
	var sliceT []string

	counter := 0
	for b.Scan() {
		counter++
		t := b.Text()
		rowSlice := strings.Split(t, " ")
		bagSlice = createBag(rowSlice)
		sliceT = append(sliceT, t)
		fmt.Println(counter)
		fmt.Println(bagSlice)

		flag := false
		for i := 0; i < len(bagSlice); i++ {
			if flag {
				canGold = append(canGold, bagSlice[i])
			}
			if  bagSlice[i].name == "shiny gold" {
				flag = true
			}
		}
		if flag {
			fmt.Println(canGold)
			fmt.Println()
		}
	}

	var bagContainGold []string
	goldBag := 0
	for _,r := range sliceT {
		for _,k := range canGold {
			if strings.Contains(r, k.name) {
				rowSlice := strings.Split(r, " ")
				bagContainGold = append(bagContainGold, rowSlice[0] + " " + rowSlice[1])
				goldBag++
				break
			}
		}
	}

	var bagContainGold2 []string
	for _,r := range sliceT {
		for _,k := range bagContainGold {
			if strings.Contains(r, k) {
				rowSlice := strings.Split(r, " ")
				bagContainGold = append(bagContainGold2, rowSlice[0] + " " + rowSlice[1])
				goldBag++
				break
			}
		}
	}



	fmt.Println(canGold)
	fmt.Println(len(canGold))
	fmt.Println(goldBag)
	fmt.Println(bagContainGold)
	fmt.Println(len(bagContainGold))
	fmt.Println(bagContainGold2)
	fmt.Println(len(bagContainGold2))
}
//[{3 clear tan} {0 dull white} {4 drab tomato} {0 clear turquoise} {2 dark coral} {0 wavy lime} {4 wavy coral} {0 striped green} {0 shiny salmon} {0 plaid gray}]

// ciao ciao 