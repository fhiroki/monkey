package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	t1 := &Boolean{Value: true}
	t2 := &Boolean{Value: true}
	f1 := &Boolean{Value: false}
	f2 := &Boolean{Value: false}

	if t1.HashKey() != t2.HashKey() {
		t.Errorf("boolean with same content have different hash keys")
	}

	if f1.HashKey() != f2.HashKey() {
		t.Errorf("boolean with same content have different hash keys")
	}

	if t1.HashKey() == f1.HashKey() {
		t.Errorf("boolean with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	o1 := &Integer{Value: 1}
	o2 := &Integer{Value: 1}
	t1 := &Integer{Value: 10}
	t2 := &Integer{Value: 10}

	if o1.HashKey() != o2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if t1.HashKey() != t2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if o1.HashKey() == t1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
