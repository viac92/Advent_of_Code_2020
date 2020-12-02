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
			if elements[i] + elements[k] == 2020 {
				fmt.Printf("%d + %d = 2020\n", elements[i], elements[k])
				multi := elements[i] * elements[k]
				fmt.Println(multi)
			}
		}
	}
}
