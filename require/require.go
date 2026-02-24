package require

import (
	"github.com/ccoveille/testifail/internal"
)

var asserter = internal.Require

// TB is the minimal interface that testifail needs to work, [*testing.T] implements it.
//
// this allows anyone to use their own testing framework as long as it implements this interface.
type TB interface {
	Helper()
	Error(args ...any)
	Fatal(args ...any)
}

// Equal asserts that expected and actual are deeply equal.
// If they are not equal, the test fails with a fatal error.
func Equal(tb TB, expected, actual any, msgAndArgs ...any) {
	asserter.Equal(tb, expected, actual, msgAndArgs...)
}

// NotEqual asserts that expected and actual are not deeply equal.
// If they are equal, the test fails with a fatal error.
func NotEqual(tb TB, expected, actual any, msgAndArgs ...any) {
	asserter.NotEqual(tb, expected, actual, msgAndArgs...)
}

// NoError asserts that err is nil.
// If err is not nil, the test fails with a fatal error.
func NoError(tb TB, err error, msgAndArgs ...any) {
	asserter.NoError(tb, err, msgAndArgs...)
}

// Error asserts that err is not nil.
// If err is nil, the test fails with a fatal error.
func Error(tb TB, err error, msgAndArgs ...any) {
	asserter.Error(tb, err, msgAndArgs...)
}

// True asserts that value is true.
// If value is false, the test fails with a fatal error.
func True(tb TB, value bool, msgAndArgs ...any) {
	asserter.True(tb, value, msgAndArgs...)
}

// False asserts that value is false.
// If value is true, the test fails with a fatal error.
func False(tb TB, value bool, msgAndArgs ...any) {
	asserter.False(tb, value, msgAndArgs...)
}

// Nil asserts that object is nil.
// If object is not nil, the test fails with a fatal error.
func Nil(tb TB, object any, msgAndArgs ...any) {
	asserter.Nil(tb, object, msgAndArgs...)
}

// NotNil asserts that object is not nil.
// If object is nil, the test fails with a fatal error.
func NotNil(tb TB, object any, msgAndArgs ...any) {
	asserter.NotNil(tb, object, msgAndArgs...)
}

// Panics asserts that calling f causes a panic.
// If f does not panic, the test fails with a fatal error.
func Panics(tb TB, f func(), msgAndArgs ...any) {
	asserter.Panics(tb, f, msgAndArgs...)
}

// Empty asserts that object is empty.
// This works with arrays, slices, maps, channels, and strings.
// If object is not empty, the test fails with a fatal error.
func Empty(tb TB, object any, msgAndArgs ...any) {
	asserter.Empty(tb, object, msgAndArgs...)
}

// Len asserts that object has the expected length.
// This works with arrays, slices, maps, channels, and strings.
// If object's length does not match, the test fails with a fatal error.
func Len(tb TB, object any, length int, msgAndArgs ...any) {
	asserter.Len(tb, object, length, msgAndArgs...)
}
