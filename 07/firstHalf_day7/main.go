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
)

type bag struct {
	num  int
	name string
}

func main() {
	b := bufio.NewScanner(os.Stdin)

	for b.Scan() {
		t := b.Text()
		fmt.Println(t)
		
	}
}
//[{3 clear tan} {0 dull white} {4 drab tomato} {0 clear turquoise} {2 dark coral} {0 wavy lime} {4 wavy coral} {0 striped green} {0 shiny salmon} {0 plaid gray}]

