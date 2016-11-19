package test

import (
	"testing"
	. "../src"
)

const (
	M = Unit_M
	CM = Unit_CM
)

func Test_should_return_true_when_equal_with_same_amount_and_same_unit(t *testing.T) {
	l1 := Length{1, M}
	l2 := Length{1, M}
	if !l1.Equal(l2) {
		t.Error()
	}
}

func Test_should_return_false_when_equal_with_same_amount_and_different_unit(t *testing.T)  {
	l1 := Length{1, M}
	l2 := Length{1, CM}
	if l1.Equal(l2) {
		t.Error()
	}
}

func Test_should_return_false_when_equal_with_different_amount_and_same_unit(t *testing.T)  {
	l1 := Length{1, M}
	l2 := Length{2, M}
	if l1.Equal(l2) {
		t.Error()
	}

}

func Test_should_return_true_when_equal_with_1M_and_100CM(t *testing.T)  {
	l1 := Length{1, M}
	l2 := Length{100, CM}
	if !l1.Equal(l2) {
		t.Error()
	}
}
