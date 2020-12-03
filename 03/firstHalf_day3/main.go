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

		for len(t) - 10 < countSteps {
			t += t
		}
		fmt.Println(len(t))
		if countRow > 0 {
			for i,r := range t {
				if r == '#' && i == countSteps {
					trees++
					break
				}
			}
		}
		countRow++
		countSteps += 3
		fmt.Println("Row ->", countRow, "Character ->", countSteps)
		fmt.Println("trees -->", trees)
	}
}