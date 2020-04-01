package lib

import (
	"errors"
	"reflect"
	"strings"
)

/*
	This retrieves a field from a struct. This supports nested structs,
	where the keys are dot-delimited (e.g., `AWS.AWSAccessKeyID`).
*/

func getField(in interface{}, field string) interface{} {
	v := reflect.ValueOf(in)
	parts := strings.Split(field, ".")
	f := v // hold this
	for i, part := range parts {

		if f.Kind() == reflect.Struct {
			// if this is a struct, grab the field by name,
			// and assign it to `f` (above)
			f = f.FieldByName(part)

			// if this field name is valid and it's the
			// last value in the `parts` array, return
			// the value
			if f.IsValid() && i == len(parts)-1 {
				return f.Interface()
			}
		} else {
			// if this isn't a struct and f isn't a valid element,
			// continue until we get something that's valid
			if !f.IsValid() {
				continue
			}

			// it's valid, let's grab the field and
			// assign it to `f` (above)
			f = f.Elem().FieldByName(part)

			// if this field name is valid and it's
			// the last value in the `parts` array,
			// return the value
			if f.IsValid() && i == len(parts)-1 {
				return f.Interface()
			}
		}
	}
	return nil
}

/*
	This sets a field in a struct. This supports nested structs,
	where the keys are dot-delimited (e.g., `AWS.AWSAccessKeyID`).
	The value can be anything, I guess.
*/
func setField(in interface{}, field string, value interface{}) (bool, error) {
	v := reflect.ValueOf(in)
	parts := strings.Split(field, ".")
	f := v // hold this
	for i, part := range parts {
		if f.Kind() == reflect.Struct {
			// if this field name is valid and it's the
			// last value in the `parts` array, set
			// the value
			if f.IsValid() && i == len(parts)-1 {
				val := reflect.ValueOf(part)
				f.Set(val)
				return true, nil
			}
		} else {
			// grab the field by name and
			// assign it to `f` (above)
			// and set the value
			f = f.Elem().FieldByName(part)
			val := reflect.ValueOf(value)

			switch f.Kind() {
			case reflect.String:
				if val.Kind() == reflect.String {
					f.SetString(val.String())
				}
			case reflect.Int:
				if val.Kind() == reflect.Int {
					f.SetInt(val.Int())
				}
			case reflect.Slice:
				if val.Kind() == reflect.Slice {
					f.Set(val)
				}
			default:
				f.Set(val)
			}
			return true, nil
		}
	}

	// this field doesn't exist dummy
	return false, errors.New("field not found:" + field)
}
