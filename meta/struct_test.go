package meta

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
	hermione := testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}

	{
		data, err := GetStructField(hermione, "born")
		if data != nil {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'born'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		exp := 1979
		data, err := GetStructField(hermione, "Born")
		if data != exp {
			t.Errorf("unexpected data: %v <=> %v", data, exp)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}

	ptr := &hermione
	{
		data, err := GetStructField(ptr, "born")
		if data != nil {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'born'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		exp := 1979
		data, err := GetStructField(ptr, "Born")
		if data != exp {
			t.Errorf("unexpected data: %v <=> %v", data, exp)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}
}

func TestSetStructField(t *testing.T) {
	hermione := testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}

	{
		data, err := SetStructField(hermione, "blood", "half-blood")
		if data != nil {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'blood'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		// Can not update field of a copyed data
		data, err := SetStructField(hermione, "Blood", "pure-blood")
		if data != nil {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "field 'Blood' can not be set" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	ptr := &hermione
	{
		data, err := SetStructField(ptr, "blood", "half-blood")
		if data != nil {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err.Error() != "no field 'blood'" {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		data, err := SetStructField(ptr, "Blood", "pure-blood")
		if data != "pure-blood" {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}
}

func TestSetStructFieldWithConvert(t *testing.T) {
	hermione := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}

	{
		data, err := SetStructField(hermione, "Born", int64(1999))
		if data != 1999 {
			t.Errorf("unexpected data: %v <=> %v", data, nil)
		}

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}

	{
		data, err := SetStructField(hermione, "Born", "pure-blood")
		if data != nil {
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

	if hermione == expected {
		t.Errorf("pointer must not be equal")
	}

	if !InstanceEqual(hermione, expected) {
		t.Errorf("unexpected data: %v <=> %v", hermione, expected)
	}

	if !reflect.DeepEqual(hermione, expected) {
		t.Errorf("unexpected data: %v <=> %v", hermione, expected)
	}
}

func TestSetStructFieldWithUntypedNil(t *testing.T) {
	hermione := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}

	data, err := SetStructField(hermione, "Wand", nil)
	if u, n := IsNil(data); !u && !n {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if !IsTypedNil(data) {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if data != hermione.Wand {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestSetStructFieldWithTypedNil(t *testing.T) {
	hermione := &testWizardType{
		Name:  "Hermione Granger",
		Born:  1979,
		Blood: "muggle-born",
		Wand: &testWandType{
			Length: 10.75,
			Core:   "dragon heartstring",
			Wood:   "vine",
		},
	}

	var wand *testWandType
	data, err := SetStructField(hermione, "Wand", wand)
	if u, n := IsNil(data); !u && !n {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if !IsTypedNil(data) {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if data != hermione.Wand {
		t.Errorf("unexpected data: %v <=> %v", data, nil)
	}

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
