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
	var inst []instruction
	var validInstruction, validInstruction2 map[int]int

	inst = append(inst, instruction{"zero", 0})
	for b.Scan() {
		t := b.Text()
		tSlice := strings.Split(t, " ")
		var instAppend instruction
		instAppend.operation = tSlice[0]
		instAppend.arggument, _ = strconv.Atoi(tSlice[1])

		inst = append(inst, instAppend)
	}

	rowCounter := 1
	accumulator := 0
	validInstruction2 = make(map[int]int)
	
	tryFirstFor := 0
	trySecondFor := 0
	
	for {
		fmt.Println("ciaoooooooooooooooooooooooooooooooooooooooooooooooo")
		
		validInstruction = make(map[int]int)
		rowCounter = 1
		accumulator = 0
		
		for rowCounter != 627 {
			var actualOperation string
			validInstruction[rowCounter]++
			fmt.Println(rowCounter)
			actualInst := inst[rowCounter]
			
			if tryFirstFor == trySecondFor {
				if actualInst.operation == "jmp" {
					actualOperation = "nop"
				}
				if actualInst.operation == "nop" {
					actualOperation = "jmp"
				}
				} else {
					actualOperation = actualInst.operation
				}
				
				fmt.Println(inst[rowCounter], actualOperation)
				
				if actualOperation == "acc" {
					accumulator += actualInst.arggument
				}
				if actualOperation == "jmp" {
					if actualInst.arggument == -1 {
						rowCounter--
					}
					if actualInst.arggument == 1 {
						rowCounter++
					}
					if actualInst.arggument < 0 {
						rowCounter += actualInst.arggument + 1
					}
					if actualInst.arggument > 0 {
						rowCounter += actualInst.arggument - 1
					}
					trySecondFor++
					continue
				}
				if actualOperation == "nop" {
					trySecondFor++
				}
				rowCounter++
				if validInstruction[rowCounter] > 1 || validInstruction2[rowCounter] > 1 {
					break
				}
			}
			validInstruction2[rowCounter] += 2
		fmt.Println(rowCounter)
		if rowCounter == 628 {
			break
		}
		trySecondFor = 0
		tryFirstFor++
	}

	fmt.Println(accumulator)
}
