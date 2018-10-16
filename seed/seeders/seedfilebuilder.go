package seeders

import (
	"encoding/csv"
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
			panic(parseError("Wrong field count on line" + string(lineCount)))
		}

		for i, v := range fieldNames {
			err := setField(newV.FieldByName(v), record[i])
			if err != nil {
				return nil, err
			}
		}

		results = append(results, newV.Interface())
		lineCount++
	}

	return results, nil
}

func setField(field reflect.Value, val string) error {
	switch t := field.Interface().(type) {
	case uint:
		v, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			return parseError("seed: Cannot convert " + val + " to uint")
		}
		field.SetUint(v)
	case uint64:
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return parseError("seed: Cannot convert " + val + " to uint64")
		}
		field.SetUint(v)
	case int:
		v, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return parseError("seed: Cannot convert " + val + " to int")
		}
		field.SetInt(v)
	case int64:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return parseError("seed: Cannot convert " + val + " to int64")
		}
		field.SetInt(v)
	case string:
		field.SetString(val)
	default:
		return parseError("seed: tried to parse a type i can't handle: " + reflect.TypeOf(t).Name())
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
