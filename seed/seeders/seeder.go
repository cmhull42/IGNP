package seeders

import "fmt"

// Seeder takes an interface for retrieving seed data. Call seed to apply this data to the db
type Seeder struct {
	mb ISeedModelBuilder
}

// NewSeeder returns a new instance
func NewSeeder(mb ISeedModelBuilder) Seeder {
	return Seeder{mb}
}

// Seed applies the seed data to the db
func (s Seeder) Seed() {
	resources, err := s.mb.ReadResources()

	if err != nil {
		// should have a db transaction here that gets rolled back
		panic(err)
	}

	fmt.Println(resources[0])
}
