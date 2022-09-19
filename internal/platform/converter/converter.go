package converter

import (
	"errors"
	"fmt"
	"reflect"
)

// Use this function if there are not fields with the same type
func FillAllByType(dataList []interface{}, model interface{}) error {
	structValue := reflect.ValueOf(model).Elem()

	for _, data := range dataList {
		for i := 0; i < structValue.NumField(); i++ {
			fill(structValue.Field(i), data)
		}
	}
	return nil
}

// Use this function if there are fields with the same type
func FillAllByFieldName(dataList map[string]interface{}, model interface{}) error {
	structValue := reflect.ValueOf(model).Elem()

	for name, data := range dataList {
		structFieldValue := structValue.FieldByName(name)
		if ok := fill(structFieldValue, data); !ok {
			return errors.New(fmt.Sprintf("The types do not match. Field name: %s. Field type: %s. Value type: %s",
				name, reflect.TypeOf(structFieldValue.Interface()), reflect.TypeOf(data)))
		}
	}

	return nil
}

func fill(structFieldValue reflect.Value, data interface{}) bool {
	fieldType := structFieldValue.Type()

	if fieldType == reflect.TypeOf(data) {
		val := reflect.ValueOf(data)
		structFieldValue.Set(val.Convert(fieldType))
		return true
	}

	return false
}
