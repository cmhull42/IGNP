package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {

	conf, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}

	var config config
	json.Unmarshal(conf, &config)

	// file://migrations is actually a folder spec. weird
	m, err := migrate.New(
		"file://migrations",
		config.DBConnString,
	)

	if err != nil {
		panic(err)
	}

	m.Down()
}

type config struct {
	DBConnString string `json:"db_conn_string"`
}
