package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func SetStructValue(kind reflect.Kind, value string, field reflect.Value, isPtr bool) error {
	if isPtr {
		uidt := reflect.TypeOf((*uuid.UUID)(nil)).Elem()
		if field.Type().Elem() == uidt {
			uid, err := uuid.Parse(value)
			if err != nil {
				return fmt.Errorf("failed to parse uuid: %w", err)
			}
			field.Set(reflect.ValueOf(&uid))
			return nil
		}

		switch field.Type().Elem().Kind() {
		case reflect.String:
			field.Set(reflect.ValueOf(&value))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ival, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse int: %w", err)
			}
			field.Set(reflect.ValueOf(&ival))

		case reflect.Float32, reflect.Float64:
			fval, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Errorf("failed to parse float: %w", err)
			}
			field.Set(reflect.ValueOf(&fval))

		case reflect.Bool:
			bval, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("failed to parse bool: %w", err)
			}
			field.Set(reflect.ValueOf(&bval))
		}

		return nil
	}

	switch kind {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ival, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse int: %w", err)
		}
		field.SetInt(ival)
	case reflect.Float32, reflect.Float64:
		fval, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("failed to parse float: %w", err)
		}
		field.SetFloat(fval)

	case reflect.Bool:
		bval, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("failed to parse bool: %w", err)
		}
		field.SetBool(bval)
	}

	return nil
}

func StructToJson(data interface{}, uid *uuid.UUID) string {
	if uid == nil {
		newUID := uuid.New()
		uid = &newUID
	}

	res, err := json.Marshal(data)
	if err != nil {
		logData := LogBuilder(*uid, "failed to marshal json", "", err)
		log.Error().Msg(logData)
	}

	return string(res)
}
