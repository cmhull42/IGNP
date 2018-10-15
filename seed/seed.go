package main

import (
	"fmt"
	"os"

	seeders "github.com/cmhull42/ignp/seed/seeders"
)

func main() {
	args := os.Args[1:]

	fmt.Print("This is a destructive operation, are you sure you want to continue? (y/n): ")
	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')
	// never change go
	//text = strings.TrimSuffix(text, "\n")
	//text = strings.TrimSuffix(text, "\r")
	//if !(text == "y" || text == "yes") {
	//	return
	//}

	operation := "--seed"
	if len(args) > 0 {
		operation = args[0]
	}

	switch operation {
	case "--seed":
		seeders.NewSeeder(seeders.CSVModelBuilder{}).Seed()
	case "--deseed":
		(seeders.Deseeder{}).Deseed()
	default:
		fmt.Println("Usage: seed {--seed|--deseed}")
	}

}
