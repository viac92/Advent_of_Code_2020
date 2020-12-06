// day 6 

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var groupSlice []string
	var findEquals, groupString string
	var yesAnswer, totalYesanswer int
	
	for b.Scan() {

		t := b.Text()
		
		for len(t) != 0 {
			groupSlice = append(groupSlice, t)
			groupString += t
			b.Scan()
			t = b.Text()
		}
		
		m := make(map[string]int)
		for id := 0; id < len(groupSlice); id++ {
			for _,r := range groupSlice[id] {
				if strings.Contains(findEquals, string(r)) == false {
					findEquals += string(r)
					m[string(r)] = 0
				}
			}
		}

		for i := 0; i < len(groupSlice); i++ {
			for _,r := range groupSlice[i] {
				_, found := m[string(r)]
				if found {
					m[string(r)]++
				}	
			}
		}

		for _, v := range m {
			if v == len(groupSlice) {
				yesAnswer++
			}
		}
		
		totalYesanswer += yesAnswer
		
		
		fmt.Println(yesAnswer)
		fmt.Println(groupSlice)
		fmt.Println(findEquals)
		fmt.Println(totalYesanswer)
		fmt.Println()
		yesAnswer = 0
		groupSlice = nil
		groupString = ""
		findEquals = ""

	}
	fmt.Println(totalYesanswer)
}