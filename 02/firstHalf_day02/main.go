//day2

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var passwordCorrette int
	var numeroPassword int

	for b.Scan() {
		t := b.Text()
		numeroPassword++
		primoNumero,_ := strconv.Atoi(t[:strings.Index(t, "-")])
		secondoNumero,_ := strconv.Atoi(t[strings.Index(t, "-") + 1:strings.Index(t, " ")])
		lettera := (t[strings.Index(t, " ") + 1:strings.Index(t, " ") + 2])
		password := (t[strings.Index(t, " ") + 4:])

		count := 0 
		for _,r := range password {
			if string(r) == lettera {
				count++
			}
		}
		if count >= primoNumero && count <= secondoNumero {
			passwordCorrette++
		}
	}
	fmt.Println(passwordCorrette, numeroPassword)
}
