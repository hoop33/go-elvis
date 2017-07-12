package elvis

import "reflect"

// Elvis returns the first argument if it's not nil, else the second argument
func Elvis(a, b interface{}) interface{} {
	// Hat tip to https://stackoverflow.com/users/187676/erik-aigner and https://stackoverflow.com/questions/13901819/quick-way-to-detect-empty-values-via-reflection-in-go
	return Ternary(a == nil || reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface()), b, a)
}

// Ternary returns a if condition is true, else b
func Ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
