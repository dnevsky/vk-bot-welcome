package pkg

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func StructToURLValues(s interface{}) (url.Values, error) {
	values := url.Values{}
	sValue := reflect.ValueOf(s)
	sType := sValue.Type()

	for i := 0; i < sValue.NumField(); i++ {
		field := sValue.Field(i)
		fieldType := sType.Field(i)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		jsonTagParts := strings.Split(jsonTag, ",")
		paramName := jsonTagParts[0]

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64: // обработка int значений структуры
			values.Set(paramName, fmt.Sprint(field.Int()))
		case reflect.String: // обработка string значений структуры
			values.Set(paramName, field.String())
		case reflect.Struct: // обработка структуры в структуре
			if paramName == "keyboard" {
				jsonData, err := json.Marshal(field.Interface())
				if err != nil {
					return nil, fmt.Errorf("failed to marshal keyboard field: %v", err)
				}
				values.Set(paramName, string(jsonData))
			} else {
				return nil, fmt.Errorf("unsupported field type: %s", field.Kind())
			}
		default: // неподдерживаемый тип данных
			return nil, fmt.Errorf("unsupported field type: %s", field.Kind())
		}
	}

	return values, nil
}
