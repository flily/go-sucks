package meta

import (
	"testing"

	"reflect"
)

func TestInterfaceEqual(t *testing.T) {
	data1 := 3
	value1 := reflect.ValueOf(data1)

	data2 := int64(3)
	value2 := reflect.ValueOf(data2)

	if value1.Interface() == value2.Interface() {
		t.Errorf("unexpected equal")
	}

	data3 := 1 + 2
	value3 := reflect.ValueOf(data3)
	if value1.Interface() != value3.Interface() {
		t.Errorf("data does not equal")
	}
}

type ttTypeCastData struct {
	Name string
	Age  int
}

func (t *ttTypeCastData) Say(word string) string {
	return t.Name + " says " + word
}

func (t ttTypeCastData) Shout(word string) string {
	return t.Name + " shouts " + word + "!"
}

type ttTypeSayingInterface interface {
	Say(string) string
}

type ttTypeShoutingInterface interface {
	Shout(string) string
}

func TestMetaTypeCasting(t *testing.T) {
	data := ttTypeCastData{
		Name: "Lily",
		Age:  33,
	}

	var infSay ttTypeSayingInterface = &data
	{
		tt := reflect.ValueOf(infSay)
		if !tt.CanConvert(reflect.TypeOf(&data)) {
			t.Errorf("%s can not convert to %s", tt.Type(), reflect.TypeOf(&data))
		}
	}

	var infShout ttTypeShoutingInterface = data
	{
		tt := reflect.ValueOf(infShout)
		if !tt.CanConvert(reflect.TypeOf(data)) {
			t.Errorf("%s can not convert to %s", tt.Type(), reflect.TypeOf(&data))
		}
	}
}
