package routes

import (
	"github.com/jmoiron/sqlx"
)

// Env is the dependency injection environment
type Env struct {
	db *sqlx.DB
}

// NewEnv returns a new enviroment with the given params
func NewEnv(db *sqlx.DB) Env {
	return Env{db}
}
