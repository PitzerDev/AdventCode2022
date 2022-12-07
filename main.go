package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	_ "strconv"
	"strings"
)

func main() {
	result := checkRanges()

	fmt.Println(result)
}

func checkRanges() (fullyEncompassingResult int) {
	var fullyEncompassingTotal int
	file, err := os.Open("./Input/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()
		currLineSplit := strings.Split(currLine, ",")
    fmt.Println(currLineSplit)
}
