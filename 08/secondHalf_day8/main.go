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
	var validInstruction map[int]int

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
	
	tryFirstFor := 0
	trySecondFor := 0
	
	for {
		fmt.Println("ciaoooooooooooooooooooooooooooooooooooooooooooooooo")
		
		validInstruction = make(map[int]int)
		rowCounter = 1
		accumulator = 0
		
		for rowCounter < 630 {
			var actualOperation string
			if tryFirstFor != trySecondFor {
				validInstruction[rowCounter]++
			}
			actualInst := inst[rowCounter]
			
			fmt.Println(rowCounter)
			fmt.Println(validInstruction)
			fmt.Println("tryFirst", tryFirstFor)
			fmt.Println("trySecond", trySecondFor)
			
			
			if tryFirstFor == trySecondFor {
				if actualInst.operation == "jmp" {
					actualOperation = "nop"
				}
				if actualInst.operation == "nop" {
					actualOperation = "jmp"
				}
				fmt.Println("swap")
				} else {
					actualOperation = actualInst.operation
				}
				if validInstruction[rowCounter] > 1 {
					break
				}
				
				fmt.Println(inst[rowCounter], actualOperation)
				
				if actualOperation == "acc" {
					accumulator += actualInst.arggument
				}
				if actualOperation == "jmp" {
					if actualInst.arggument < 0 {
						rowCounter += actualInst.arggument
					}
					if actualInst.arggument > 0 {
						rowCounter += actualInst.arggument
					}
					trySecondFor++
					continue
				}
				if actualOperation == "nop" {
					trySecondFor++
				}
				rowCounter++
			}
		fmt.Println(rowCounter)
		if rowCounter >= 630 {
			break
		}
		trySecondFor = 0
		tryFirstFor++
	}

	fmt.Println(accumulator)
}
