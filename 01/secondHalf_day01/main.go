//day one

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var elements []int

	for b.Scan(){
		t, _ := strconv.Atoi(b.Text())
		elements = append(elements, t)
	}

	fmt.Println(len(elements))

	for i := 0; i < len(elements); i++ {
		for k := 0; k < len(elements); k++ {
			for j := 0; j < len(elements); j++ {
				if elements[i] + elements[k] + elements[j] == 2020 {
					fmt.Printf("%d + %d + %d = 2020\n", elements[i], elements[k], elements[j])
					multi := elements[i] * elements[k] * elements[j]
					fmt.Println(multi)
				}
			}
		}
	}
}