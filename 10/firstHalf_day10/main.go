package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"sort"
)

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var adapters []int
	var recursion map[int]int

	b := bufio.NewScanner(file)

	for b.Scan() {
		t := b.Text()
		num, _ := strconv.Atoi(t)
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)
	recursion = make(map[int]int)
	currentJ := 0
	for _,adaps := range adapters {
		differrance := adaps - currentJ
		currentJ = adaps
		recursion[differrance]++	
	}
	recursion[3]++
	
	fmt.Printf("recursion of 1 -> %d\nrecursion of 3 -> %d\n", recursion[1],recursion[3])
	fmt.Printf("Answer (rec 1 * rec 3) = %d\n", recursion[3]*recursion[1])
}