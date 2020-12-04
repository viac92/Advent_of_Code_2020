// day 4

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
	"strconv"
	"unicode"
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
	var passaportiValidi [][]string
	var contaSbagliati, passaportiGiusti int
	
	var controllaPassaporti int

	for b.Scan() {
		t := b.Text()
		
		for len(t) != 0 {
			idStringa += t + " "
			b.Scan()
			t = b.Text()
		}
		sliceOrdinata := strings.Split(idStringa, " ")
		sliceOrdinata = sliceOrdinata[: len(sliceOrdinata) - 1]
		fmt.Println(sliceOrdinata)
		for i := 0; i < len(sliceOrdinata); i++ {
				tag := sliceOrdinata[i]
				tag = tag[:3]
				tagSlice = append(tagSlice, tag)
		}
		sort.Strings(tagSlice)

		if len(tagSlice) <= 6 {
			contaSbagliati++
		} else if len(tagSlice) == 7 && contains(tagSlice, "cid") {
			contaSbagliati++
		} else {
			passaportiValidi = append(passaportiValidi, sliceOrdinata)	
		}
		idStringa = ""
		tagSlice = nil
	}

	for i := 0; i < len(passaportiValidi); i++ {
		for j := 0; j < len(passaportiValidi[i]); j++ {
			fmt.Println(passaportiValidi[i][j])
			switch passaportiValidi[i][j][:3] {
			case "byr":
				annoByr, _ := strconv.Atoi(passaportiValidi[i][j][4:])
				if  annoByr >= 1920 &&  annoByr <= 2002 {
					controllaPassaporti++
				}
			case "iyr":
				annoIyr, _ := strconv.Atoi(passaportiValidi[i][j][4:])
				if annoIyr >= 2010 && annoIyr <= 2020 {
					controllaPassaporti++
				}
			case "eyr":
				annoEyr, _ := strconv.Atoi(passaportiValidi[i][j][4:])
				if annoEyr >= 2020 && annoEyr <= 2030 {
					controllaPassaporti++
				}
			case "hgt":
				switch string(passaportiValidi[i][j][len(passaportiValidi[i][j]) - 2:]) {
				case "cm":
					if len(passaportiValidi[i][j][4:len(passaportiValidi[i][j]) - 2]) != 3 {
						break
					}
					cm, ok := strconv.Atoi(passaportiValidi[i][j][4:len(passaportiValidi[i][j]) - 2])
					if ok != nil {
						break
					}
					if cm < 150 || cm > 193 {
						break
					}
					controllaPassaporti++
				case "in":
					if len(passaportiValidi[i][j][4:len(passaportiValidi[i][j]) - 2]) != 2 {
						break
					}
					inch, ok := strconv.Atoi(passaportiValidi[i][j][4:len(passaportiValidi[i][j]) - 2])
					if ok != nil {
						break
					}
					if inch < 59 || inch > 76 {
						break
					}
					controllaPassaporti++
				}
			case "hcl":
				if string(passaportiValidi[i][j][4]) == "#" && len(passaportiValidi[i][j][4:]) == 7 {
					for _, r := range passaportiValidi[i][j][4:] {
						if strings.IndexRune("0987654321abcdef", r) == -1 {
							break
							} 
						}
						controllaPassaporti++
					}
				case "ecl":
					switch passaportiValidi[i][j][4:] {
					case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
						controllaPassaporti++
					}
				case "pid":
					digit := 0
					for _, r := range passaportiValidi[i][j][4:] {
						if unicode.IsDigit(r) {
							digit++
						}
					}
					if digit == 9 {
					controllaPassaporti++
				}
			}
		}
		fmt.Println(controllaPassaporti)
		if controllaPassaporti == 7 {
			passaportiGiusti++
		}
		controllaPassaporti = 0
	}
	fmt.Println(passaportiGiusti)
}
