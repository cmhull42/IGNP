package seeders

import (
	"database/sql"
	"strings"

	sysmodel "github.com/cmhull42/ignp/model/system"
	_ "github.com/go-sql-driver/mysql" // "justify this blank import" i dunno man, that's what the docs say
)

// Seeder takes an interface for retrieving seed data. Call seed to apply this data to the db
type Seeder struct {
	mb ISeedModelBuilder
	c  Config
}

// NewSeeder returns a new instance
func NewSeeder(mb ISeedModelBuilder, c Config) Seeder {
	return Seeder{mb, c}
}

// Seed applies the seed data to the db
func (s Seeder) Seed() (err error) {

	// TODO: handle invalid connection string in a sane way instead of assuming it's correct
	parts := strings.Split(s.c.ConnectionString, "://")
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

	// make all seed calls in a transaction, having a half seeded database doesn't make sense
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// does the fun stuff, actually puts the data in the tables using the transaction
	err = populateTables(tx, s)
	if err != nil {
		// ruh roh
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func populateTables(tx *sql.Tx, s Seeder) (err error) {
	var resources []sysmodel.Resource
	var locations []sysmodel.Location
	var resourceLocations []sysmodel.ResourceLocation
	var resourceTypes []sysmodel.ResourceType

	if resourceTypes, err = s.mb.ReadResourceTypes(); err != nil {
		return err
	}

	for _, rt := range resourceTypes {
		if _, err := tx.Exec("insert into SystemResourceTypes (id,name) values (?,?)",
			rt.ResourceTypeID,
			rt.Name,
		); err != nil {
			return err
		}
	}

	if resources, err = s.mb.ReadResources(); err != nil {
		return err
	}

	for _, r := range resources {
		if _, err := tx.Exec("insert into SystemResources (id,resourcetype,name,rarity) values (?, ?, ?, ?)",
			r.ResourceID,
			r.ResourceType,
			r.Name,
			r.Rarity,
		); err != nil {
			return err
		}
	}

	if locations, err = s.mb.ReadLocations(); err != nil {
		return err
	}

	for _, l := range locations {
		if _, err := tx.Exec("insert into SystemLocations (id,coordx,coordy) values (?, ?)",
			l.LocationID,
			l.CoordX,
			l.CoordY,
		); err != nil {
			return err
		}
	}

	if resourceLocations, err = s.mb.ReadResourceLocations(); err != nil {
		return err
	}

	for _, rl := range resourceLocations {
		if _, err := tx.Exec("insert into GenResourceLocations (resourcetype,locationid,capacity) values (?, ?, ?)",
			rl.ResourceType,
			rl.LocationID,
			rl.Capacity,
		); err != nil {
			return err
		}
	}

	return nil
}
