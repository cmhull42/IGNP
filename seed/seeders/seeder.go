package seeders

import (
	"fmt"

	sysmodel "github.com/cmhull42/ignp/model/system"
)

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
	var resources []sysmodel.Resource
	var locations []sysmodel.Location
	var resourceLocations []sysmodel.ResourceLocation
	var resourceTypes []sysmodel.ResourceType
	var err error

	resources, err = s.mb.ReadResources()

	if err != nil {
		// should have a db transaction here that gets rolled back
		panic(err)
	}

	locations, err = s.mb.ReadLocations()

	if err != nil {
		panic(err)
	}

	resourceLocations, err = s.mb.ReadResourceLocations()

	if err != nil {
		panic(err)
	}

	resourceTypes, err = s.mb.ReadResourceTypes()

	fmt.Println(resources)
	fmt.Println(locations)
	fmt.Println(resourceLocations)
	fmt.Println(resourceTypes)
}
