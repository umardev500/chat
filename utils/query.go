package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// BuildUpdateQuery builds update query based on struct
//
// Parameters:
//
//   - tbl: table name
//   - s: pointer of struct
//   - condition: the query condition
//   - startIndex: let say we give it 1 it will create params start from $2...$n
func BuildUpdateQuery(tbl string, s interface{}, condition string, startIndex int) (string, []interface{}, error) {
	vType := reflect.TypeOf(s)
	if vType.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("build update query expected struct, got %T", s)
	}

	var fieldNames []string
	var args []interface{}

	v := reflect.ValueOf(s)
	vNum := v.NumField()

	for i := 0; i < vNum; i++ {
		field := vType.Field(i)
		fieldValue := v.Field(i).Interface()
		fieldName, ok := field.Tag.Lookup("db")
		if !ok || fieldName == "-" || reflect.ValueOf(fieldValue).IsZero() {
			continue
		}

		fieldNames = append(fieldNames, fmt.Sprintf("%s = $%d", fieldName, len(args)+1+startIndex))
		args = append(args, fieldValue)
	}

	setClause := strings.Join(fieldNames, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tbl, setClause, condition)

	return query, args, nil
}

func BuildBulkInsertQuery(tbl string, s interface{}) (string, []interface{}, error) {
	vType := reflect.TypeOf(s)
	vNum := vType.Elem().NumField()
	v := reflect.ValueOf(s)
	vLen := v.Len()
	if vType.Kind() != reflect.Slice {
		return "", nil, fmt.Errorf("build bulk insert query expected slice, got %T", s)
	}

	var fieldNames []string
	for i := 0; i < vNum; i++ {
		field := vType.Elem().Field(i)
		fieldName, ok := field.Tag.Lookup("db")
		if !ok || fieldName == "-" {
			continue
		}

		fieldNames = append(fieldNames, fieldName)
	}

	var argsParams []string
	var args []interface{}
	for i := 0; i < vLen; i++ {
		sv := v.Index(i)
		// svType := sv.Type()
		svNum := sv.NumField()

		var placeholders []string
		for j := 0; j < svNum; j++ {
			fv := sv.Field(j)
			placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)+1))
			if fv.Kind() == reflect.Slice {
				fieldValueJson, err := json.Marshal(fv.Interface())
				if err != nil {
					return "", nil, err
				}
				args = append(args, string(fieldValueJson))
				continue
			}
			args = append(args, fv.Interface())

		}
		argsParams = append(argsParams, fmt.Sprintf("(%s)", strings.Join(placeholders, ", ")))
	}

	columns := strings.Join(fieldNames, ",")
	argsValue := strings.Join(argsParams, ",")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", tbl, columns, argsValue)

	return query, args, nil
}
