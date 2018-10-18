package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	seeders "github.com/cmhull42/ignp/seed/seeders"
)

type config struct {
	DBConnString string `json:"db_conn_string"`
}

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

	fmt.Println()

	conf, err := ioutil.ReadFile("../conf.json")
	if err != nil {
		panic(err)
	}

	var config config
	if err := json.Unmarshal(conf, &config); err != nil {
		panic(err)
	}

	operation := "--seed"
	if len(args) > 0 {
		operation = args[0]
	}

	switch operation {
	case "--seed":
		if err := seeders.NewSeeder(
			seeders.CSVModelBuilder{},
			seeders.Config{ConnectionString: config.DBConnString},
		).Seed(); err != nil {
			panic(err)
		}
	case "--deseed":
		if err := seeders.NewDeseeder(
			seeders.Config{ConnectionString: config.DBConnString},
		).Deseed(); err != nil {
			panic(err)
		}
	default:
		fmt.Println("Usage: seed {--seed|--deseed}")
	}

}
