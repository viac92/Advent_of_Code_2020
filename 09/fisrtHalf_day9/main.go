// day 9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rowsCounter int
	var premble []int
	var checkSum map[int]int
	var answer string
	b := bufio.NewScanner(os.Stdin)



	rowsCounter = 1
	for b.Scan() {
		t := b.Text()
		if rowsCounter > 26 {
			num,_ := strconv.Atoi(t)
			if checkSum[num] == 1 {
				premble = append(premble[1:], num)
			} else {
				answer = t
				break
			}
		}


		if rowsCounter <= 25 {
			num,_ := strconv.Atoi(t)
			premble = append(premble, num)
		} else {
			checkSum = make(map[int]int)
			for i,num := range premble {
				fmt.Println(num)
				for _,num2 := range premble[1+i:] {
					fmt.Println(num2)
					if checkSum[num+num2] < 1 {
						checkSum[num+num2]++
					}
				}
			}
			fmt.Println(premble)
			fmt.Println(checkSum)
			
		}
		rowsCounter++
	}

	fmt.Println(answer)

}