package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Println(input)
}
