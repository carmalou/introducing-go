package math

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})

	if v != 1.5 {
		t.Error("expected 1.5. got:", v)
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{1, 2, 3, 4})

	if v != 1 {
		t.Error("expected 1. got:", v)
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{1, 2, 3, 4})

	if v != 4 {
		t.Error("expected 1. got:", v)
	}
}
