package tests

import (
	"reflect"
	"sort"
	"testing"

	"github.com/braydenandrson/collections/collections"
)

func TestNewCounterFromSlice(t *testing.T) {
	c := NewCounter[string]([]string{"a", "b", "a", "c", "b", "a"})

	expected := map[string]int{"a": 3, "b": 2, "c": 1}
	if !reflect.DeepEqual(c.Items(), expected) {
		t.Errorf("Expected %v, got %v", expected, c.Items())
	}
}

func TestNewCounterFromMap(t *testing.T) {
	c := NewCounter[string](map[string]int{"a": 2, "b": 3})

	if c.Get("a") != 2 || c.Get("b") != 3 {
		t.Errorf("Incorrect values for keys a or b")
	}
	if c.Get("c") != 0 {
		t.Errorf("Expected zero for missing key c")
	}
}

func TestAddAndAddN(t *testing.T) {
	c := NewCounter[string]()
	c.Add("x")
	c.AddN("x", 4)

	if c.Get("x") != 5 {
		t.Errorf("Expected 5, got %d", c.Get("x"))
	}
}

func TestMostCommon(t *testing.T) {
	c := NewCounter[string]([]string{"apple", "banana", "apple", "orange", "banana", "apple"})
	common := c.MostCommon(2)

	if len(common) != 2 || common[0].Item != "apple" || common[0].Count != 3 {
		t.Errorf("Unexpected most common result: %v", common)
	}
}

func TestElements(t *testing.T) {
	c := NewCounter[string](map[string]int{"a": 2, "b": 1, "c": 0, "d": -2})
	result := c.Elements()
	sort.Strings(result)

	expected := []string{"a", "a", "b"}
	sort.Strings(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestUpdateAndSubtract(t *testing.T) {
	c := NewCounter[string](map[string]int{"a": 3, "b": 2})
	c.Update(map[string]int{"a": 2, "c": 4})
	c.Subtract(map[string]int{"b": 1, "c": 2})

	expected := map[string]int{"a": 5, "b": 1, "c": 2}
	if !reflect.DeepEqual(c.Items(), expected) {
		t.Errorf("Expected %v, got %v", expected, c.Items())
	}
}

func TestMathOperations(t *testing.T) {
	c := NewCounter[string](map[string]int{"a": 3, "b": 1})
	d := NewCounter[string](map[string]int{"a": 1, "b": 2})

	sum := c.AddCounter(d)
	if sum.Get("a") != 4 || sum.Get("b") != 3 {
		t.Errorf("Addition failed: %v", sum.Items())
	}

	diff := c.SubCounter(d)
	if !reflect.DeepEqual(diff.Items(), map[string]int{"a": 2}) {
		t.Errorf("Subtraction failed: %v", diff.Items())
	}

	and := c.AndCounter(d)
	if !reflect.DeepEqual(and.Items(), map[string]int{"a": 1, "b": 1}) {
		t.Errorf("Intersection failed: %v", and.Items())
	}

	or := c.OrCounter(d)
	if !reflect.DeepEqual(or.Items(), map[string]int{"a": 3, "b": 2}) {
		t.Errorf("Union failed: %v", or.Items())
	}
}

func TestComparisons(t *testing.T) {
	c := NewCounter[string](map[string]int{"a": 2, "b": 0})
	d := NewCounter[string](map[string]int{"a": 2})

	if !c.Equals(d) {
		t.Errorf("Counters should be equal (zero-counts ignored)")
	}
	if !d.IsSubsetOf(c) || !c.IsSupersetOf(d) {
		t.Errorf("Subset/superset failed")
	}
}

func TestUnary(t *testing.T) {
	c := NewCounter[string](map[string]int{"a": 2, "b": -3, "c": 0})

	pos := c.Positive()
	if !reflect.DeepEqual(pos.Items(), map[string]int{"a": 2}) {
		t.Errorf("Positive failed: %v", pos.Items())
	}

	neg := c.Negative()
	if !reflect.DeepEqual(neg.Items(), map[string]int{"b": 3}) {
		t.Errorf("Negative failed: %v", neg.Items())
	}
}
