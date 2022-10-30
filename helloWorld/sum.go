package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	cleanInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return cleanInput
}

func main() {
	a := getInput()
	b := getInput()
	fmt.Println("The sum of ", a, "and", b, "is", a+b)
}
