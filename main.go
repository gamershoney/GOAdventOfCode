package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string = "read.txt"

var cv int = 0
var counter int = 0

func main() {
	items, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	buf := bytes.Lines(items)

	for line := range buf {
		str := strings.Split(string(line), "")
		dir := str[0]

		num := strings.Join(str[1:len(str)], "")
		val, _ := strconv.Atoi(num)
		if dir == "R" {
			if (cv + val) > 99 {
				cv = cv + val - 100
			} else {
				cv = cv + val
			}
		} else {
			if val > cv {
				fmt.Println("value:", val, "Is greater than current value:", cv)
				val = val - cv
				cv = 100 - val
			} else {
				cv = cv - val
			}
		}

		if cv == 0 {
			counter++
		}

	}

	err = os.WriteFile("output.txt", items, 0777)
	if err != nil {
		fmt.Println(err)
	}

}
