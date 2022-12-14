package meta

import (
	"testing"

	"net"
	"reflect"

	"github.com/flily/go-sucks/meta/reflectutil"
)

func TestDuplicateForSimpleValue(t *testing.T) {
	{
		data := 1
		got, err := Duplicate(data)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !InstanceEqual(data, got) {
			t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
				got, reflect.TypeOf(got), data, reflect.TypeOf(data))
		}

		if !Equal(data, got.(int)) {
			t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
				got, reflect.TypeOf(got), data, reflect.TypeOf(data))
		}
	}
}

func TestDuplicateForSimpleStruct(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	data := testStruct{
		Name: "Harry Potter",
		Age:  13,
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithPrivateField(t *testing.T) {
	type testStruct struct {
		Name string
		age  int
	}

	data := testStruct{
		Name: "Harry Potter",
		age:  13,
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithPrivateArray(t *testing.T) {
	type testStruct struct {
		Name    string
		courses []string
	}

	data := testStruct{
		Name: "Harry Potter",
		courses: []string{
			"Potions",
			"Charms",
		},
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	data.courses = append(data.courses, "Astronomy")
	if InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithPrivatePointerArray(t *testing.T) {
	type testStruct struct {
		Name    string
		age     int
		parents []*testStruct
	}

	data := testStruct{
		Name: "Harry Potter",
		age:  13,
		parents: []*testStruct{
			{"James Potter", 33, nil},
			{"Lily Potter", 33, nil},
		},
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	data.parents[0].Name = "Sirius Black"
	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithPrivatePointerVariable(t *testing.T) {
	type testStruct struct {
		Name   string
		age    int
		spouse *testStruct
	}

	ginny := &testStruct{
		Name: "Ginny Weasley",
		age:  12,
	}

	data := testStruct{
		Name:   "Harry Potter",
		age:    13,
		spouse: ginny,
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got.(testStruct).spouse, reflect.TypeOf(got), data.spouse, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	data.spouse.Name = "Cho Chang"
	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithPrivatePointerLiteral(t *testing.T) {
	type testStruct struct {
		Name   string
		age    int
		spouse *testStruct
	}

	data := testStruct{
		Name: "Harry Potter",
		age:  13,
		spouse: &testStruct{
			Name: "Ginny Weasley",
			age:  12,
		},
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	data.spouse.Name = "Cho Chang"
	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithPrivateMap(t *testing.T) {
	type testStruct struct {
		Name   string
		scores map[string]int
	}

	data := testStruct{
		Name: "Harry Potter",
		scores: map[string]int{
			"Potions": 98,
			"Charms":  95,
		},
	}

	got, err := Duplicate(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	if !Equal(data, got.(testStruct)) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}

	data.scores["Astronomy"] = 96
	if InstanceEqual(data, got) {
		t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
			got, reflect.TypeOf(got), data, reflect.TypeOf(data))
	}
}

func TestDuplicateForSimpleStructWithChannelField(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
		Ch   chan int
	}

	data := testStruct{
		Name: "Harry Potter",
		Age:  13,
		Ch:   make(chan int),
	}

	got, err := Duplicate(data)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if err.Error() != "kind chan is not duplicatable" {
		t.Errorf("unexpected error: %v", err)
	}

	if got != nil {
		t.Errorf("expected nil, but got %#v (%s)", got, reflect.TypeOf(got))
	}
}

func TestDuplicateNil(t *testing.T) {
	{
		var n *int

		got, err := Duplicate(n)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if got != n {
			t.Errorf("unexpected result: %#v (%s) <=> %#v (%s)",
				got, reflect.TypeOf(got), n, reflect.TypeOf(n))
		}
	}

	{
		var n *interface{}

		got, err := Duplicate(n)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if got != n {
			t.Errorf("unexpected result: %#v (%T) <=> %#v (%T)",
				got, got, n, n)
		}
	}

	{
		got, err := Duplicate(nil)
		if err != ErrUntypedNil {
			t.Errorf("unexpected error: %v", err)
		}

		if got != nil {
			t.Errorf("unexpected result: %#v (%T)", got, got)
		}
	}
}

func TestDuplicateNilValue(t *testing.T) {
	nilValue := reflect.ValueOf(nil)
	copy, err := DuplicateValueInstance(nilValue)
	if err == nil {
		t.Errorf("expected ErrUntypedNil, got nil")
	}

	if copy.IsValid() {
		t.Errorf("unexpected result: %#v (%T)", copy, copy)
	}

	if u, n := reflectutil.NilType(copy); !u || n {
		t.Errorf("unexpected result: %#v (%T)", copy, copy)
	}
}

func TestDuplicateNormalData(t *testing.T) {
	dataList := []interface{}{
		0,
		42,
		0.0,
		uint(42),
		3.1415926,
		"",
		"lorem ipsum",
		true,
		false,
	}

	for _, x := range dataList {
		d, err := Duplicate(x)
		if err != nil {
			t.Errorf("unexpected error on %v: %v", x, err)
		}

		if x != d {
			t.Errorf("value not equal %v <=> %v", x, d)
		}
	}
}

func TestDuplicateStructData(t *testing.T) {
	type testData struct {
		I int
		S string
		F float32
		B bool
		N interface{}
	}

	s := testData{
		I: 42,
		S: "lorem ipsum",
		F: 3.1415926,
		B: true,
		N: nil,
	}

	c := s
	if c != s {
		t.Errorf("value not equal %v <=> %v", s, c)
	}

	d, err := Duplicate(s)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if d != s {
		t.Errorf("value not equal %v <=> %v", s, d)
	}

	p, err := Duplicate(&s)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if pt, ok := (p.(*testData)); ok {
		if *pt != s {
			t.Errorf("unexpected result: %T (%#v)", p, p)
		}

	} else {
		t.Errorf("value not equal %v <=> %v", s, pt)
	}
}

func TestDuplicateArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b, err := Duplicate(a)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	c := b.([]int)
	for i := range a {
		if a[i] != c[i] {
			t.Errorf("value not equal %v <=> %v", a, b)
		}
	}

	a[3] = 42
	for i := range a {
		if i == 3 {
			if c[3] == a[3] {
				t.Errorf("value not equal %v <=> %v", a, b)
			}
			continue
		}

		if a[i] != c[i] {
			t.Errorf("value not equal %v <=> %v", a, b)
		}
	}
}

func TestDuplicatePointerArray(t *testing.T) {
	ptr := func(n int) *int { return &n }
	{
		a := []*int{ptr(1), ptr(2), ptr(3), ptr(4), nil}

		b, err := Duplicate(a)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		c := b.([]*int)
		for i := range a {
			if a[i] != c[i] {
				t.Errorf("value not equal %v <=> %v", a, b)
			}
		}
	}
}

func TestDuplicateArrayError(t *testing.T) {
	c := make(chan int, 1)
	a := []interface{}{1, c}

	got, err := Duplicate(a)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if got != nil {
		t.Errorf("unexpected result: %#v (%T)", got, got)
	}
}

func TestDuplicateMapError(t *testing.T) {
	ch := make(chan int, 1)

	{
		m := map[chan int]int{
			ch: 42,
		}

		got, err := Duplicate(m)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if err.Error() != "kind chan is not duplicatable" {
			t.Errorf("unexpected error: %v", err)
		}

		if got != nil {
			t.Errorf("unexpected result: %#v (%T)", got, got)
		}
	}

	{
		m := map[int]chan int{
			42: ch,
		}

		got, err := Duplicate(m)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if err.Error() != "kind chan is not duplicatable" {
			t.Errorf("unexpected error: %v", err)
		}

		if got != nil {
			t.Errorf("unexpected result: %#v (%T)", got, got)
		}
	}
}

func TestDuplicateChannel(t *testing.T) {
	c := make(chan int, 1)
	c <- 1

	got, err := Duplicate(c)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if got != nil {
		t.Errorf("unexpected result: %#v", got)
	}
}

func TestDuplicateStructInterface(t *testing.T) {
	// This case do not go into duplicateInterface
	raw := &net.AddrError{
		Err:  "test",
		Addr: "lorem ipsum",
	}
	var err error = raw
	ep := &err

	got, err := Duplicate(ep)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	gotErr, ok := got.(*error)
	if !ok {
		t.Fatalf("unexpected result: %#v (%T)", got, got)
	}

	gotInfo, ok := (*gotErr).(*net.AddrError)
	if !ok {
		t.Fatalf("unexpected result: %#v (%T)", got, got)
	}

	if (*gotInfo).Addr != raw.Addr || (*gotInfo).Err != raw.Err {
		t.Errorf("unexpected result: %#v (%T) <=> %#v (%T)",
			gotInfo, gotInfo, raw, raw)
	}
}

func TestDuplicateAnyInterface(t *testing.T) {
	a := []interface{}{
		nil,
		1,
		1.2,
		&net.AddrError{
			Err: "test",
		},
	}

	got, err := Duplicate(a)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !InstanceEqual(a, got) {
		t.Errorf("unexpected result: %#v (%T) <=> %#v (%T)",
			got, got, a, a)
	}
}
