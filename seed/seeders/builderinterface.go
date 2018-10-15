package seeders

import sysmodel "github.com/cmhull42/ignp/model/system"

// ISeedModelBuilder returns populated models of data to be seeded
type ISeedModelBuilder interface {
	ReadResources() ([]sysmodel.Resource, error)
}
