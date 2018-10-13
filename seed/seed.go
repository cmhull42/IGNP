package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	fmt.Print("This is a destructive operation, are you sure you want to continue? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// never change go
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	if !(text == "y" || text == "yes") {
		return
	}

	operation := "--seed"
	if len(args) > 0 {
		operation = args[0]
	}

	switch operation {
	case "--seed":
		seed()
	case "--deseed":
		deseed()
	default:
		fmt.Println("Usage: seed {--seed|--deseed}")
	}

}

func seed() {
	fmt.Println("j/k i didn't do anything")
}

func deseed() {
	fmt.Println("doing nothing has been undone")
}
