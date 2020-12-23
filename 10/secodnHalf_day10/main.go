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
	
	for _,i := range adapters {
		fmt.Println(i)
	}
	
	recursion := make(map[int]int)
	currentJ := 0

	for i := 0; i <  len(adapters);  i++ {
		fmt.Printf("i = %d - %d\n", i, adapters[i])
		
		differrance := adapters[i] - currentJ
		currentJ = adapters[i]
		recursion[differrance]++
		
		if i == len(adapters) - 6 {
			break
		}
		
	}
}