package statsp

import (
	"reflect"
	"testing"
)

func TestCleanerGuages(t *testing.T) {
	cleaner := NewCleaner()
	tests := []struct {
		relative bool
		in       float64
		out      float64
	}{
		{true, 1.0, 1.0},
		{true, 1.0, 2.0},
		{true, -1.0, 1.0},
		{false, 1.0, 1.0},
		{false, -1.0, -1.0},
		{true, -1.0, 0.0},
	}
	for i, tst := range tests {
		metric := Metric{"foo", Guage, tst.relative, tst.in, 0}
		expected := Metric{"foo", Guage, false, tst.out, 0}
		cleaned := cleaner.Clean(metric)
		if cleaned != expected {
			t.Errorf("%d: expected %v got %v", i, expected, cleaned)
		}
	}
}

func TestCleanerOthers(t *testing.T) {
	cleaner := NewCleaner()
	tests := []Metric{
		Metric{"foo", Counter, true, -1.0, 0},
		Metric{"foo", Counter, true, 1.0, 0},
		Metric{"foo", Timer, true, 1.0, 0},
		Metric{"foo", Set, true, -1.0, 0},
	}
	for i, tst := range tests {
		cleaned := cleaner.Clean(tst)
		expected := tst
		expected.Relative = false
		if cleaned != expected {
			t.Errorf("%d: expected %v got %v", i, expected, cleaned)
		}
	}
}

func TestCleanMetrics(t *testing.T) {
	in := []Metric{
		Metric{"boo", Counter, true, -1.0, 0},
		Metric{"foo", Guage, true, 1.0, 0},
		Metric{"foo", Guage, true, 1.0, 0},
		Metric{"foo", Guage, true, -1.0, 0},
		Metric{"foo", Guage, false, 1.0, 0},
	}
	expected := []Metric{
		Metric{"boo", Counter, false, -1.0, 0},
		Metric{"foo", Guage, false, 1.0, 0},
		Metric{"foo", Guage, false, 2.0, 0},
		Metric{"foo", Guage, false, 1.0, 0},
		Metric{"foo", Guage, false, 1.0, 0},
	}
	cleaner := NewCleaner()
	cleaned := cleaner.CleanMetrics(in)
	if !reflect.DeepEqual(cleaned, expected) {
		t.Errorf("expected\n%v\ngot\n%v\n", expected, cleaned)
	}
}
