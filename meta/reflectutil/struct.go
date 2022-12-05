package reflectutil

import (
	"reflect"
)

func GetStructFieldValue(data interface{}, field string) (reflect.Value, error) {
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}

	var err error
	fieldValue := dataValue.FieldByName(field)
	if !fieldValue.IsValid() {
		err = ErrReflect.Derive("no field '%s'", field)
	}

	return fieldValue, err
}

func GetStructField(data interface{}, field string) (interface{}, error) {
	fieldValue, err := GetStructFieldValue(data, field)
	if err != nil {
		return nil, err
	}

	return fieldValue.Interface(), nil
}

func SetStructFieldValue(data interface{}, field string, value reflect.Value) (interface{}, error) {
	fieldValue, err := GetStructFieldValue(data, field)
	if err != nil {
		return nil, err
	}

	if !fieldValue.CanSet() {
		return nil, ErrReflect.Derive("field '%s' can not be set", field)
	}

	untypedNil, typedNil := NilType(value)
	if untypedNil {
		newNil := NewTypedNil(fieldValue.Type())
		fieldValue.Set(newNil)

	} else if typedNil {
		fieldValue.Set(NewTypedNil(fieldValue.Type()))

	} else {
		fieldValueType := fieldValue.Type()
		if value.Type() != fieldValueType {
			if value.CanConvert(fieldValueType) {
				convertedValue := value.Convert(fieldValueType)
				fieldValue.Set(convertedValue)

			} else {
				return nil, ErrReflect.Derive("field '%s' requires type %s, but %s",
					field, fieldValue.Type(), value.Type())
			}

		} else {
			fieldValue.Set(value)
		}
	}

	return fieldValue.Interface(), nil
}

func SetStructField(data interface{}, field string, value interface{}) (interface{}, error) {
	valueValue := reflect.ValueOf(value)
	return SetStructFieldValue(data, field, valueValue)
}
