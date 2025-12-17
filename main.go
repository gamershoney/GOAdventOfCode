package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var file string = "read.txt"

var cv int = 50
var counter int = 0
var prev int

func main() {
	reg := regexp.MustCompile(`\D`)
	items, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	buf := bytes.Lines(items)
	cL := 0
	for line := range buf {
		str := strings.Split(string(line), "")
		dir := str[0]

		num := strings.Join(str[1:len(str)], "")
		num = reg.ReplaceAllString(num, "")
		val, err := strconv.Atoi(num)
		if val >= 100 {
			val = val % 100
		}
		if val == 0 {
			break
		}

		if dir == "R" {
			if (cv + val) > 99 {
				cv = cv + val - 100
			} else {
				cv = cv + val
			}
		} else {
			if val > cv {
				val = val - cv
				cv = 100 - val
			} else {
				cv = cv - val
			}
		}
		if err != nil {
			fmt.Println("error: ", err)
		}
		if (prev == 0) && (cv == 0) {
			break
		}

		if cv == 0 {
			counter++
		}

		prev = cv
		cL++
		fmt.Println("Read ", cL, "Lines")
	}
	fmt.Println("Total 0's hit:", counter)
	err = os.WriteFile("output.txt", items, 0777)
	if err != nil {
		fmt.Println(err)
	}

}
