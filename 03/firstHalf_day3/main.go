// day3

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var countSteps, countRow, trees int

	for b.Scan() {
		t := b.Text()

		for len(t) < countSteps {
			t += t
		}

		if countRow > 0 {
			for i,r := range t {
				if r == '#' && i == countSteps {
					trees++
				}
			}
		}
		countRow++
		countSteps += 3
		fmt.Println("Riga ->", countRow, "\nlettera ->", countSteps)
		fmt.Println(trees)
	}
}