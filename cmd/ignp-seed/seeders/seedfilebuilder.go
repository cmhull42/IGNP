package seeders

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	sysmodel "github.com/cmhull42/ignp/model/system"
)

const locationPath string = "data/system/locations.csv"
const resourcePath string = "data/system/resources.csv"
const resourceLocationPath string = "data/system/resourcelocations.csv"
const resourceTypesPath string = "data/system/resourcetypes.csv"

// parseError is a local error type used to handle bad csv data
type parseError string

// implementation of the Error interface
func (p parseError) Error() string {
	return string(p)
}

// CSVModelBuilder is an implementation of ISeedModelBuilder that populates the model from csv files
type CSVModelBuilder struct {
}

// ReadResources implements ISeedModelBuilder.ReadResources
func (cmb CSVModelBuilder) ReadResources() ([]sysmodel.Resource, error) {
	r, err := readType(resourcePath, &sysmodel.Resource{})
	if err != nil {
		return nil, err
	}

	resources := make([]sysmodel.Resource, len(r))
	for i := range r {
		resources[i] = r[i].(sysmodel.Resource)
	}
	return resources, nil
}

// ReadLocations implements ISeedModelBuilder.ReadLocations
func (cmb CSVModelBuilder) ReadLocations() ([]sysmodel.Location, error) {
	r, err := readType(locationPath, &sysmodel.Location{})
	if err != nil {
		return nil, err
	}

	locations := make([]sysmodel.Location, len(r))
	for i := range r {
		locations[i] = r[i].(sysmodel.Location)
	}
	return locations, nil
}

// ReadResourceLocations implements ISeedModelBuilder.ReadResourceLocations
func (cmb CSVModelBuilder) ReadResourceLocations() ([]sysmodel.ResourceLocation, error) {
	r, err := readType(resourceLocationPath, &sysmodel.ResourceLocation{})
	if err != nil {
		return nil, err
	}

	resourceLocations := make([]sysmodel.ResourceLocation, len(r))
	for i := range r {
		resourceLocations[i] = r[i].(sysmodel.ResourceLocation)
	}
	return resourceLocations, nil
}

// ReadResourceTypes implements ISeedModelBuilder.ReadResourceTypes
func (cmb CSVModelBuilder) ReadResourceTypes() ([]sysmodel.ResourceType, error) {
	r, err := readType(resourceTypesPath, &sysmodel.ResourceType{})
	if err != nil {
		return nil, err
	}

	resourceTypes := make([]sysmodel.ResourceType, len(r))
	for i := range r {
		resourceTypes[i] = r[i].(sysmodel.ResourceType)
	}
	return resourceTypes, nil
}

func readType(parsePath string, parseType interface{}) (res []interface{}, perr error) {
	defer func() {
		if r := recover(); r != nil {
			res = nil
			perr = r.(parseError)
		}
	}()

	f, err := os.Open(parsePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	results := make([]interface{}, 0)

	r := csv.NewReader(f)

	fieldNames, err := r.Read()
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(parseType).Elem()

	lineCount := 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		newV := reflect.New(t).Elem()

		if len(record) != t.NumField() {
			panic(parseError("wrong field count on line" + string(lineCount)))
		}

		for i, v := range fieldNames {
			if setField(newV.FieldByName(v), record[i]) != nil {
				return nil, fmt.Errorf("seed: cannot convert value from string to %s: %v", newV.Type().Name(), err)
			}
		}

		results = append(results, newV.Interface())
		lineCount++
	}

	return results, nil
}

func setField(field reflect.Value, val string) (err error) {
	switch field.Type().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var v uint64
		v, err = strconv.ParseUint(val, 10, field.Type().Bits())
		field.SetUint(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var v int64
		v, err = strconv.ParseInt(val, 10, 0)
		field.SetInt(v)
	case reflect.String:
		field.SetString(val)
	case reflect.Float32, reflect.Float64:
		var v float64
		v, err = strconv.ParseFloat(val, field.Type().Bits())
		field.SetFloat(v)
	default:
		panic(parseError("seed: tried to parse a type i can't handle: " + field.Type().Name()))
	}

	return nil
}

// generics when
func indexOf(s []string, val string) int {
	for p, v := range s {
		if v == val {
			return p
		}
	}
	return -1
}
