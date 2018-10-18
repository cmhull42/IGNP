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
		if err := m.Up(); err != nil {
			panic(err)
		}
		return
	}

	if args[0] == "--down" {
		if err := m.Down(); err != nil {
			panic(err)
		}
	} else if args[0] == "--up" {
		if err := m.Up(); err != nil {
			panic(err)
		}
	} else if args[0] == "--version" {
		i, _ := strconv.ParseUint(args[1], 10, 0)
		if err := m.Migrate(uint(i)); err != nil {
			panic(err)
		}
	} else if args[0] == "--force" {
		i, _ := strconv.ParseInt(args[1], 10, 0)
		if err := m.Force(int(i)); err != nil {
			panic(err)
		}
	} else {
		if err := m.Up(); err != nil {
			panic(err)
		}
		return
	}
}
