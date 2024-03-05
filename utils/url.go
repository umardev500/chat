package utils

import (
	"reflect"

	"github.com/rs/zerolog/log"
)

func ParseQueryString(data map[string]string, s interface{}) (err error) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr {
		log.Error().Msgf("not a pointer")
		return
	}
	num := val.Elem().NumField()

	for i := 0; i < num; i++ {
		field := val.Elem().Type().Field(i)

		// if the field is a struct
		// then parse the query to it
		// case for nested struct
		if field.Type.Kind() == reflect.Struct {
			fieldValue := val.Elem().Field(i)

			err = ParseQueryString(data, fieldValue.Addr().Interface())
			if err != nil {
				return err
			}

			continue
		}

		paramName, ok := field.Tag.Lookup("query")
		if !ok {
			continue
		}

		// Get the value
		paramValue := data[paramName]

		fieldVal := val.Elem().Field(i)
		if paramValue != "" {
			switch fieldVal.Kind() {
			case reflect.Ptr:
				err = SetStructValue(fieldVal.Kind(), paramValue, fieldVal, true)
				if err != nil {
					return err
				}
			default:
				err = SetStructValue(fieldVal.Kind(), paramValue, fieldVal, false)
				if err != nil {
					return err
				}
			}
		} else {
			// if the field has a default value
			defaultValue, ok := field.Tag.Lookup("default")
			if !ok && defaultValue != "" {
				continue
			}

			fieldVal := val.Elem().Field(i)
			switch fieldVal.Kind() {
			case reflect.Ptr:
				err = SetStructValue(fieldVal.Kind(), defaultValue, fieldVal, true)
				if err != nil {
					return err
				}
			default:
				err = SetStructValue(fieldVal.Kind(), defaultValue, fieldVal, false)
				if err != nil {
					return err
				}
			}
		}
	}

	return
}
