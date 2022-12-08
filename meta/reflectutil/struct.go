package reflectutil

import (
	"reflect"
)

func GetStructField(data reflect.Value, field string) (reflect.Value, error) {
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}

	if data.Kind() != reflect.Struct {
		return reflect.Value{}, ErrReflect.Derive("data is not a struct: %s", data.Kind())
	}

	var err error
	fieldValue := data.FieldByName(field)
	if !fieldValue.IsValid() {
		err = ErrReflect.Derive("no field '%s'", field)
	}

	return fieldValue, err
}

func SetStructField(data reflect.Value, field string, value reflect.Value) (reflect.Value, error) {
	fieldValue, err := GetStructField(data, field)
	if err != nil {
		return reflect.Value{}, err
	}

	if !fieldValue.CanSet() {
		return reflect.Value{}, ErrReflect.Derive("field '%s' can not be set", field)
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
				return reflect.Value{}, ErrReflect.Derive("field '%s' requires type %s, but %s",
					field, fieldValue.Type(), value.Type())
			}

		} else {
			fieldValue.Set(value)
		}
	}

	return fieldValue, nil
}
