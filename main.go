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

var cv int = 0
var counter int = 0

func main() {
	reg := regexp.MustCompile(`\D`)
	items, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	buf := bytes.Lines(items)

	for line := range buf {
		str := strings.Split(string(line), "")
		dir := str[0]

		num := strings.Join(str[1:len(str)], "")
		num = reg.ReplaceAllString(num, "")
		val, _ := strconv.Atoi(num)
		if dir == "R" {
			cv = cv + val
		} else {
			cv = cv - val
		}

		if cv == 0 {
			counter++
			fmt.Println(counter)
		}

	}

	err = os.WriteFile("output.txt", items, 0777)
	if err != nil {
		fmt.Println(err)
	}

}
