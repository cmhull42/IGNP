package seeders

import (
	"database/sql"
	"strings"
)

// Deseeder reverses the seed operation leaving the database clean
type Deseeder struct {
	c Config
}

// NewDeseeder returns a new Deseeder with the passed config
func NewDeseeder(c Config) Deseeder {
	return Deseeder{c}
}

// Deseed starts the deseeding operation
func (d Deseeder) Deseed() error {
	// TODO: handle invalid connection string in a sane way instead of assuming it's correct
	parts := strings.Split(d.c.ConnectionString, "://")
	driver, dataSourceName := parts[0], parts[1]

	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return err
	}

	defer db.Close()

	// test the database connection
	if err := db.Ping(); err != nil {
		return err
	}

	// "exceptions aren't elegant" - go devs
	if _, err := db.Exec("delete from GenResourceLocations"); err != nil {
		return err
	}
	if _, err := db.Exec("delete from SystemLocations"); err != nil {
		return err
	}
	if _, err := db.Exec("delete from SystemResources"); err != nil {
		return err
	}
	if _, err := db.Exec("delete from SystemResourceTypes"); err != nil {
		return err
	}

	return nil
}
