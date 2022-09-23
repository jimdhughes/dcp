package main

import "testing"

func TestIsMaxValueAndKeys(t *testing.T) {
	m := MagicStruct{}
	m.Init()
	m.Plus("a")
	m.Plus("a")
	m.Plus("b")
	if !m.IsMaxValue(2) {
		t.Fail()
	}
	if m.GetMax() != "a" {
		t.Fail()
	}
}

func TestIsMinValueAndKeys(t *testing.T) {
	m := MagicStruct{}
	m.Init()
	m.Plus("a")
	m.Plus("a")
	m.Plus("b")
	if !m.IsMinValue(1) {
		t.Fail()
	}
	if !m.IsKeyMinValue("b") {
		t.Errorf("Expected min key value to be b. Min value was: %s", m.MinValues[len(m.MinValues)].Key)
	}
	if m.IsKeyMinValue("a") {
		t.Error("a came back as a min value key", m.MinValues[len(m.MinValues)].Key)
	}
}

func TestOneKeyWasMaxThenWeRemovedItMakingAnotherMax(t *testing.T) {
	m := MagicStruct{}
	m.Init()
	m.Plus("a")
	m.Plus("a")
	m.Plus("b")
	m.Plus("b")
	m.Plus("b")
	m.Minus("b")
	if !m.IsKeyMaxValue("b") {
		t.Error("Expect b to still be max as was greater or equal to previous entry")
	}
	m.Minus("b")
	if m.IsKeyMaxValue("b") {
		t.Error("Expected a to take over as the max sized ey now")
	}
	if m.IsKeyMaxValue("a") {
		t.Log("a was the largets key as expected")
	}
}

func TestOneKeyWasMinThenWeAddedThingsAndRemovedThings(t *testing.T) {
	m := MagicStruct{}
	m.Init()
	m.Plus("a")
	if !m.IsKeyMinValue("a") {
		t.Errorf("key a should be both min and max after first entry")
	}
	m.Plus("b")
	m.Plus("a")
	m.Minus("a")
	m.Minus("b")
	if m.GetMin() != "a" {
		t.Errorf("key a should be min value again")
	}
	m.Plus("a")
	m.Plus("a")
	m.Plus("b")
	if m.GetMin() != "b" {
		t.Errorf("b should be the min value now")
	}
}

func TestMinKeyIsNoLongerThereWhenThereShouldBeNoEntriesLeft(t *testing.T) {
	m := MagicStruct{}
	m.Init()
	m.Plus("a")
	m.Plus("b")
	m.Plus("a")
	m.Minus("b")
	if m.GetMin() != "a" {
		t.Errorf("Expected a to be min value. got %s", m.GetMin())
	}
}

func TestBaseCase(t *testing.T) {
	m := MagicStruct{}
	m.Init()
	m.Plus("a")
	m.Plus("a")
	m.Plus("b")
	m.Plus("b")
	m.Plus("b")
	m.Plus("a")
	m.Minus("b")
	m.Plus("c")
	if m.GetMax() != "a" {
		t.Errorf("A should be the max. got %s", m.GetMax())
	}
	if m.GetMin() != "c" {
		t.Errorf("c should be the min. got %s", m.GetMin())
	}
}
