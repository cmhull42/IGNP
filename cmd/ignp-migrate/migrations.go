package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type config struct {
	DBConnString string `json:"db_conn_string"`
}

func main() {

	var confFile string
	var up bool
	var down bool
	var version int
	var force int
	flag.StringVar(&confFile, "conf", "", "path to the config file")
	flag.BoolVar(&up, "up", false, "migrate up to the latest version")
	flag.BoolVar(&down, "down", false, "migrate down to the oldest version")
	flag.IntVar(&version, "version", -1, "migrate to a specific version")
	flag.IntVar(&force, "force", -1, "same as version, but forces if the db state is unclean")

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
	json.Unmarshal(conf, &config)

	// file://migrations is actually a folder spec. weird
	m, err := migrate.New(
		"file://scripts",
		config.DBConnString,
	)

	if err != nil {
		panic(err)
	}

	if down {
		if err := m.Down(); err != nil {
			panic(err)
		}
	} else if up {
		if err := m.Up(); err != nil {
			panic(err)
		}
	} else if version != -1 {
		if err := m.Migrate(uint(version)); err != nil {
			panic(err)
		}
	} else if force != -1 {
		if err := m.Force(force); err != nil {
			panic(err)
		}
	} else {
		if err := m.Up(); err != nil {
			panic(err)
		}
		return
	}
}
