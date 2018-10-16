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

const resourcePath string = "data/system/resources.csv"

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
