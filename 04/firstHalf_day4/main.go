// day 4

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
)

func contains(s []string, str string) bool {
	for _,v := range s {
		if v == str {
			return true
		}
	}
	return false
}



func main() {
	b := bufio.NewScanner(os.Stdin)
	var idStringa string
	var tagSlice []string
	var contaSbagliati, contaGiusti int
	

	for b.Scan() {
		t := b.Text()
		
		for len(t) != 0 {
			idStringa += t + " "
			b.Scan()
			t = b.Text()
		}
		//fmt.Println(idStringa)
		sliceOrdinata := strings.Split(idStringa, " ")
		for i := 0; i < len(sliceOrdinata) - 1; i++ {
				tag := sliceOrdinata[i]
				tag = tag[:3]
				tagSlice = append(tagSlice, tag)
		}
		sort.Strings(tagSlice)
		fmt.Println(tagSlice)

		if len(tagSlice) <= 6 {
			fmt.Println("Sbagliato")
			contaSbagliati++
		} else if len(tagSlice) == 7 && contains(tagSlice, "cid") {
			fmt.Println("sbagliato")
			contaSbagliati++
		} else {
			fmt.Println("Giusto")
			contaGiusti++
		}
		
		idStringa = ""
		tagSlice = nil
	}

	fmt.Println(contaGiusti, contaSbagliati)
}
