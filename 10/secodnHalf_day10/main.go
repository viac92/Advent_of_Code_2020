package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func main() {
	file,err := os.Open("puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var adapters []int

	b := bufio.NewScanner(file)
	for b.Scan() {
		t := b.Text()
		num,_ := strconv.Atoi(t)
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1] + 3)
	fmt.Println(adapters)


	recursion := make(map[int]int)
	currentJ := 0

	option := 1
	for i := 0; i <  len(adapters);  i++ {
		differrance := adapters[i] - currentJ
		currentJ = adapters[i]
		recursion[differrance]++

		if i == len(adapters) - 1 {
			break
		}

		fmt.Println(adapters[i])

		if adapters[i+1] - adapters[i] == 1 && adapters[i+2] - adapters[i] == 2 {
			option++
			fmt.Println("uno",option)
			if adapters[i+3] - adapters[i] == 3 {
				option++
				fmt.Println("due",option)
			}
		}
	}

	fmt.Println(recursion)
	fmt.Println(option)
}