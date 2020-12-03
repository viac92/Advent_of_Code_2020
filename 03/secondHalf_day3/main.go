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
		if countRow == 0 {
			countSteps++
		}
		countRow++
		if countRow % 2 == 1 && countRow != 1 {
			for i,r := range t {
				if r == '#' && i == countSteps {
					trees++
					break
				}
			}
			fmt.Println("Row ->", countRow, "Character ->", countSteps)
			fmt.Println("trees -->", trees)
			countSteps++
		}
	}
	
}


//R1, D1 = 53
//R3, D1 = 167
//R5, D1 = 54
//R7, D1 = 67
//R1, D2 = 23