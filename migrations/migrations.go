package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type config struct {
	DBConnString string `json:"db_conn_string"`
}

func main() {

	args := os.Args[1:]

	conf, err := ioutil.ReadFile("../conf.json")
	if err != nil {
		panic(err)
	}

	var config config
	json.Unmarshal(conf, &config)

	// file://migrations is actually a folder spec. weird
	m, err := migrate.New(
		"file://scripts",
		config.DBConnString,
	)

	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		m.Up()
		return
	}

	if args[0] == "--down" {
		m.Down()
	} else if args[0] == "--up" {
		m.Up()
	} else if args[0] == "--version" {
		i, _ := strconv.ParseUint(args[1], 10, 0)
		m.Migrate(uint(i))
	} else {
		m.Up()
		return
	}
}
