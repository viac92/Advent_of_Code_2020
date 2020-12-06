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
		
		for id := 0; id < len(groupSlice); id++ {
			for _,r := range groupSlice[id] {
				if strings.Contains(findEquals, string(r)) == false {
					findEquals += string(r)
					yesAnswer++
				}
			}
			fmt.Println(yesAnswer)
			totalYesanswer += yesAnswer
			yesAnswer = 0
		}

		totalYesanswer += yesAnswer

		fmt.Println(groupSlice)
		fmt.Println(findEquals)
		fmt.Println(totalYesanswer)
		fmt.Println()
		groupSlice = nil
		groupString = ""
		findEquals = ""

	}
	fmt.Println(totalYesanswer)
}