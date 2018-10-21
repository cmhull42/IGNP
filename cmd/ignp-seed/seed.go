package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	seeders "github.com/cmhull42/ignp/cmd/ignp-seed/seeders"
)

type config struct {
	DBConnString string `json:"db_conn_string"`
}

func indexOf(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func main() {
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

	var confFile string
	var seed bool
	var deseed bool
	flag.StringVar(&confFile, "conf", "", "path to the config file")
	flag.BoolVar(&seed, "seed", false, "use this flag if you'd like to seed new data")
	flag.BoolVar(&deseed, "deseed", false, "use this flag if you'd like to clear seeded data")

	flag.Parse()

	if confFile == "" {
		flag.PrintDefaults()
		return
	}

	conf, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}

	var config config
	if err := json.Unmarshal(conf, &config); err != nil {
		panic(err)
	}

	if !deseed {
		if err := seeders.NewSeeder(
			seeders.CSVModelBuilder{},
			seeders.Config{ConnectionString: config.DBConnString},
		).Seed(); err != nil {
			panic(err)
		}
	} else {
		if err := seeders.NewDeseeder(
			seeders.Config{ConnectionString: config.DBConnString},
		).Deseed(); err != nil {
			panic(err)
		}
	}

}

func printUsage() {
	fmt.Println("Usage: seed --conf <conf> {--seed|--deseed}")
}
