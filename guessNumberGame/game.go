package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	cleanInput, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return cleanInput
}

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("Guess The Number :)")
	for x := 10; x >= 0; x-- {
		fmt.Print("Enter Number: ")
		number := getInput()
		if number == target {
			fmt.Println("Ahan! you caught me")
			break
		} else if number > target {
			fmt.Println("Oops! your guess was High. You have ", x-1, "tries remaining")
		} else {
			fmt.Println("Oops! your guess was Low. You have ", x-1, "tries remaining")
		}
	}

}
