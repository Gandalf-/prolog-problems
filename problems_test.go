package problems

/*
	https://sites.google.com/site/prologsite/prolog-problems
	Gandalf- 2019
*/

import (
	"testing"
)

func TestLastElement(t *testing.T) {
	a := sequence(10)

	if v, _ := LastElement(a); v != 9 {
		t.Error("a was incorrect")
	}

	b := sequence(0)
	if _, err := LastElement(b); err == nil {
		t.Error("no error")
	}
}

func TestSecondToLast(t *testing.T) {
	a := sequence(10)

	if v, _ := SecondToLast(a); v != 8 {
		t.Error("a was incorrect")
	}

	b := sequence(0)
	if _, err := SecondToLast(b); err == nil {
		t.Error("no error")
	}

	c := sequence(1)
	if _, err := SecondToLast(c); err == nil {
		t.Error("no error")
	}
}

func TestKthElement(t *testing.T) {

	a := sequence(10)

	if _, err := KthElement(a, -1); err == nil {
		t.Error("no error")
	}

	if _, err := KthElement(a, 10); err == nil {
		t.Error("no error")
	}

	if v, _ := KthElement(a, 9); v != 9 {
		t.Error("no error")
	}
}

func TestCount(t *testing.T) {

	a := sequence(10)
	b := sequence(0)

	if Count(a) != 10 {
		t.Error("wrong count")
	}

	if Count(b) != 0 {
		t.Error("wrong count")
	}
}

func TestReverse(t *testing.T) {

	a := sequence(10)
	b := sequence(0)

	ra := Reverse(a)
	rb := Reverse(b)

	if equal(a, ra) {
		t.Error("a was not changed")
	}

	if !equal(b, rb) {
		t.Error("b was changed")
	}

	for i := range ra {
		if ra[i] != 9-i {
			t.Errorf("%d != %d", ra[i], 9-i)
		}
	}
}
