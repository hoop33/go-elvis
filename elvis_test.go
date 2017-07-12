package elvis

import "testing"

var elvisTests = []struct {
	a        interface{}
	b        interface{}
	expected interface{}
}{
	{"Hello", nil, "Hello"},
	{nil, "Hi", "Hi"},
	{"Hello", "Hi", "Hello"},
	{"", "Hi", "Hi"},
	{0, 1, 1},
	{1, 2, 1},
	{0, 1.5, 1.5},
}

var ternaryTests = []struct {
	condition bool
	a         interface{}
	b         interface{}
	expected  interface{}
}{
	{true, "Hello", "Hi", "Hello"},
	{false, "Hello", "Hi", "Hi"},
	{true, 0, 1, 0},
	{false, 0, 1, 1},
	{true, 2.3, 3.3, 2.3},
	{false, 2.3, 3.3, 3.3},
}

func TestElvis(t *testing.T) {
	for _, tt := range elvisTests {
		actual := Elvis(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Elvis(%v, %v): expected %v, actual %v", tt.a, tt.b, tt.expected, actual)
		}
	}
}

func TestTernary(t *testing.T) {
	for _, tt := range ternaryTests {
		actual := Ternary(tt.condition, tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Ternary(%t, %v, %v): expected %v, actual %v", tt.condition, tt.a, tt.b, tt.expected, actual)
		}
	}
}
