// day8

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	arggument int
}

func main() {
	b := bufio.NewScanner(os.Stdin)
	var inst[] instruction
	var validInstruction map[int]int

	inst = append(inst, instruction{"zero", 0})
	for b.Scan() {
		t := b.Text()
		tSlice := strings.Split(t," ")
		var instAppend instruction
		instAppend.operation = tSlice[0]
		instAppend.arggument,_ = strconv.Atoi(tSlice[1])

		inst = append(inst, instAppend)
	}

	rowCounter := 1
	accumulator := 0
	validInstruction = make(map[int]int)
	for rowCounter <= 628 {
		validInstruction[rowCounter]++
		if validInstruction[rowCounter] > 1 {
			break
		}
		fmt.Println(rowCounter)
		fmt.Println(inst[rowCounter])
		actualInst := inst[rowCounter]
		
		if actualInst.operation == "acc" {
			accumulator += actualInst.arggument
		} 
		if actualInst.operation == "jmp" {
			if actualInst.arggument < 0 {
				rowCounter += actualInst.arggument
			} 
			if actualInst.arggument > 0 {
				rowCounter += actualInst.arggument
			}
			continue
		}
	
		rowCounter++
	}
	
	

	fmt.Println(accumulator)
}