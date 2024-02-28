package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ParseData struct {
	Field   string  `json:"field"`
	Tag     string  `json:"tag"`
	Params  *string `json:"params,omitempty"`
	Message string  `json:"message"`
}

func parseTag(tag, field, param string) ParseData {
	var response ParseData
	var paramValue *string
	if param != "" {
		paramValue = &param
	}

	switch tag {
	case "required":
		msg := fmt.Sprintf("%s is required", field)
		response = ParseData{
			Field:   field,
			Tag:     tag,
			Params:  paramValue,
			Message: msg,
		}
	case "gte":
		msg := fmt.Sprintf("%s must be greater than or equal to params", field)
		response = ParseData{
			Field:   field,
			Tag:     tag,
			Params:  paramValue,
			Message: msg,
		}
	case "oneof":
		params := strings.Split(param, " ")
		paramStr := strings.Join(params, ",")

		msg := fmt.Sprintf("%s must one of params", field)
		response = ParseData{
			Field:   field,
			Tag:     tag,
			Params:  &paramStr,
			Message: msg,
		}
	}

	return response
}

func ValidateStruct(v *validator.Validate, s interface{}) ([]ParseData, error) {
	var responseData []ParseData

	if err := v.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			tag := err.Tag()
			param := err.Param()

			var tempField reflect.StructField
			tempField, ok := reflect.TypeOf(s).FieldByName(field)
			if !ok {

				// Handle nested struct
				sv := reflect.ValueOf(s)
				for j := 0; j < sv.NumField(); j++ {
					// Get the field's type
					fieldType := sv.Type().Field(j)
					// Check if the field's type is a struct
					if fieldType.Type.Kind() == reflect.Struct {
						resp, _ := ValidateStruct(v, sv.Field(j).Interface())
						responseData = append(responseData, resp...)
					}
				}

			}

			// Lookup json tag
			fJsonName, ok := tempField.Tag.Lookup("json")
			if !ok {
				continue
			}

			// append json name if found
			data := parseTag(tag, fJsonName, param)
			responseData = append(responseData, data)
		}

		return responseData, err
	}

	return nil, nil
}
