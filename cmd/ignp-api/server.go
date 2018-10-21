package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/cmhull42/ignp/cmd/ignp-api/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type config struct {
	DBConnString string `json:"db_conn_string"`
}

func buildRoutes(db *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	e := routes.NewEnv(db)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/resources", routes.ResourceRoutes(e))
	})

	return router
}

func main() {

	var confFile string
	flag.StringVar(&confFile, "conf", "", "path of the config file")

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

	// TODO: handle invalid connection string in a sane way instead of assuming it's correct
	parts := strings.Split(config.DBConnString, "://")
	driver, dataSourceName := parts[0], parts[1]

	db, err := sqlx.Connect(driver, dataSourceName)
	if err != nil {
		panic(err)
	}

	router := buildRoutes(db)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":9804", router))
}
