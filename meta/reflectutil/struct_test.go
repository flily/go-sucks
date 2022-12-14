package reflectutil

import (
	"testing"

	"reflect"
)

type testWandType struct {
	Length float32
	Core   string
	Wood   string
}

type testWizardType struct {
	Name  string
	Born  int
	Blood string
	Wand  *testWandType
}

func TestGetStructField(t *testing.T) {
	testData := testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}
	hermione := reflect.ValueOf(testData)

	{
		data, err := GetStructField(hermione, "born")
		if data.Kind() != reflect.Invalid {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'born'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		exp := 1979
		data, err := GetStructField(hermione, "Born")
		if data.Interface() != exp {
			t.Errorf("unexpected data: %v <=> %v", data, exp)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}

	ptr := reflect.ValueOf(&testData)
	{
		data, err := GetStructField(ptr, "born")
		if data.Kind() != reflect.Invalid {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'born'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		exp := 1979
		data, err := GetStructField(ptr, "Born")
		if data.Interface() != exp {
			t.Errorf("unexpected data: %v <=> %v", data, exp)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}
}

func TestGetStructFieldOnInteger(t *testing.T) {
	data, err := GetStructField(reflect.ValueOf(42), "born")
	if data.Kind() != reflect.Invalid {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if err.Error() != "data is not a struct: int" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetStructField(t *testing.T) {
	testData := testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}
	hermione := reflect.ValueOf(testData)

	{
		data, err := SetStructField(hermione, "blood", reflect.ValueOf("half-blood"))
		if data.Kind() != reflect.Invalid {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'blood'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		// Can not update field of a copyed data
		data, err := SetStructField(hermione, "Blood", reflect.ValueOf("pure-blood"))
		if data.Kind() != reflect.Invalid {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "field 'Blood' can not be set" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	ptr := reflect.ValueOf(&testData)
	{
		data, err := SetStructField(ptr, "blood", reflect.ValueOf("half-blood"))
		if data.Kind() != reflect.Invalid {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'blood'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		exp := "pure-blood"
		data, err := SetStructField(ptr, "Blood", reflect.ValueOf("pure-blood"))
		if data.Interface() != exp {
			t.Errorf("unexpected data: %v <=> %v", data, exp)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}
}

func TestSetStructFieldWithConvert(t *testing.T) {
	testData := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}
	hermione := reflect.ValueOf(testData)

	{
		data, err := SetStructField(hermione, "Born", reflect.ValueOf(int64(1999)))
		if data.Interface() != 1999 {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		data, err := SetStructField(hermione, "Born", reflect.ValueOf("pure-blood"))
		if data.Kind() != reflect.Invalid {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "field 'Born' requires type int, but string" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	expected := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1999,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}

	if testData == expected {
		t.Errorf("pointer must not be equal")
	}

	if !reflect.DeepEqual(testData, expected) {
		t.Errorf("unexpected data: %v <=> %v", hermione, expected)
	}
}

func TestSetStructFieldWithUntypedNil(t *testing.T) {
	testData := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}
	hermione := reflect.ValueOf(testData)

	data, err := SetStructField(hermione, "Wand", reflect.ValueOf(nil))
	if data.Interface() != (*testWandType)(nil) {
		t.Errorf("unexpected data: %v (%T) <=> %v", data, data, nil)
	}

	if data.Interface() != testData.Wand {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetStructFieldWithTypedNil(t *testing.T) {
	testData := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}
	hermione := reflect.ValueOf(testData)

	var wand *testWandType
	data, err := SetStructField(hermione, "Wand", reflect.ValueOf(wand))
	if data.Interface() != (*testWandType)(nil) {
		t.Errorf("unexpected data: %v (%T) <=> %v", data, data, nil)
	}

	if data.Interface() != testData.Wand {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
