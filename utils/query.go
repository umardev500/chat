package utils

import (
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
