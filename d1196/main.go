package main

type KeyValuePairs struct {
	Key   string
	Value int
}

type MagicStruct struct {
	Values    map[string]int
	MaxValues []KeyValuePairs
	MinValues []KeyValuePairs
}

func (m *MagicStruct) Init() {
	m.Values = make(map[string]int)
	m.MaxValues = make([]KeyValuePairs, 0)
	m.MinValues = make([]KeyValuePairs, 0)
}

func (m *MagicStruct) IsMaxValue(value int) bool {
	if len(m.MaxValues) == 0 {
		return true
	}
	return m.MaxValues[len(m.MaxValues)-1].Value <= value
}

func (m *MagicStruct) IsMinValue(value int) bool {
	if len(m.MinValues) == 0 {
		return true
	}
	return m.MinValues[len(m.MinValues)-1].Value >= value
}

func (m *MagicStruct) IsKeyMaxValue(key string) bool {
	if len(m.MaxValues) == 0 {
		return true
	}
	maxVal := m.MaxValues[len(m.MaxValues)-1]
	return maxVal.Key == key
}

func (m *MagicStruct) IsKeyMinValue(key string) bool {
	if len(m.MinValues) == 0 {
		return true
	}
	minVal := m.MinValues[len(m.MinValues)-1]
	return minVal.Key == key
}

func (m *MagicStruct) Plus(key string) {
	val := m.Values[key]
	val = val + 1
	m.Values[key] = val
	// handle isMinValue
	if m.IsMinValue(val) {
		m.MinValues = append(m.MinValues, KeyValuePairs{Key: key, Value: val})
	}
	// handle isMaxValue
	if m.IsMaxValue(val) {
		m.MaxValues = append(m.MaxValues, KeyValuePairs{Key: key, Value: val})
	}
}

func (m *MagicStruct) Minus(key string) {
	val := m.Values[key]
	if val <= 0 {
		delete(m.Values, key)
	} else {
		m.Values[key] = val - 1
	}
	if m.IsKeyMinValue(key) {
		m.MinValues = m.MinValues[0 : len(m.MinValues)-1]
	}
	if m.IsKeyMaxValue(key) {
		m.MaxValues = m.MaxValues[0 : len(m.MaxValues)-1]
	}
}

func (m *MagicStruct) GetMax() string {
	if len(m.MaxValues) == 0 {
		return ""
	}
	return m.MaxValues[len(m.MaxValues)-1].Key
}

func (m *MagicStruct) GetMin() string {
	if len(m.MinValues) == 0 {
		return ""
	}
	return m.MinValues[len(m.MinValues)-1].Key
}
