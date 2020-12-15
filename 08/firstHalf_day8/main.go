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

	for b.Scan() {
		t := b.Text()
		tSlice := strings.Split(t," ")
		var instAppend instruction
		instAppend.operation = tSlice[0]
		instAppend.arggument,_ = strconv.Atoi(tSlice[1])

		inst = append(inst, instAppend)
	}

	rowCounter := 0
	accumulator := 0
	for rowCounter < 628 {
		actualInst := inst[rowCounter]
		
		if actualInst.operation == "acc" {
			accumulator += actualInst.arggument
		} 
		if actualInst.operation == "jmp" {
			rowCounter += actualInst.arggument
			continue
		}
	
		fmt.Println(rowCounter+1)
		rowCounter++
	}
	
	

	fmt.Println(accumulator)
}