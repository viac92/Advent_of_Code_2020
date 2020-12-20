// day 9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var num []int

	for b.Scan() {
		t := b.Text()
		
		n,_ := strconv.Atoi(t)
		num = append(num, n)
	}

	sum := 0
	highN := 0
	weakness := 0
	for i,n := range num {
		sum = n
		fmt.Println("nuovo")
		for j := i + 1; j < len(num); j++ {
			sum += num[j]
			fmt.Printf("Numero piccolo - %d\nNumero grande - %d\nSomma - %d\n\n",n, num[j], sum)
			if sum == 10884537 {
				highN = num[j]
				break
			}
			if sum > 10884537 {
				break
			}
		}
		if sum == 10884537 {
			fmt.Println(n, highN)
			weakness = n + highN
			break
		}
	}

	fmt.Println(sum)
	fmt.Println(weakness)
}