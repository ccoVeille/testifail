package assert

import (
	"github.com/ccoveille/testifail/internal"
)

var asserter = internal.Assert

// TB is the minimal interface that testifail needs to work, [*testing.T] implements it.
//
// this allows anyone to use their own testing framework as long as it implements this interface.
type TB interface {
	Helper()
	Error(args ...any)
	Fatal(args ...any)
}

// Equal reports whether expected and actual are deeply equal.
// It returns true on success.
func Equal(tb TB, expected, actual any, msgAndArgs ...any) bool {
	return asserter.Equal(tb, expected, actual, msgAndArgs...)
}

// NotEqual reports whether expected and actual are not deeply equal.
// It returns true on success.
func NotEqual(tb TB, expected, actual any, msgAndArgs ...any) bool {
	return asserter.NotEqual(tb, expected, actual, msgAndArgs...)
}

// NoError reports whether err is nil.
// It returns true on success.
func NoError(tb TB, err error, msgAndArgs ...any) bool {
	return asserter.NoError(tb, err, msgAndArgs...)
}

// Error reports whether err is non-nil.
// It returns true on success.
func Error(tb TB, err error, msgAndArgs ...any) bool {
	return asserter.Error(tb, err, msgAndArgs...)
}

// True reports whether value is true.
// It returns true on success.
func True(tb TB, value bool, msgAndArgs ...any) bool {
	return asserter.True(tb, value, msgAndArgs...)
}

// False reports whether value is false.
// It returns true on success.
func False(tb TB, value bool, msgAndArgs ...any) bool {
	return asserter.False(tb, value, msgAndArgs...)
}

// Nil reports whether object is nil, including typed nils.
// It returns true on success.
func Nil(tb TB, object any, msgAndArgs ...any) bool {
	return asserter.Nil(tb, object, msgAndArgs...)
}

// NotNil reports whether object is not nil.
// It returns true on success.
func NotNil(tb TB, object any, msgAndArgs ...any) bool {
	return asserter.NotNil(tb, object, msgAndArgs...)
}

// Panics reports whether f panics when invoked.
// It returns true on success.
func Panics(tb TB, f func(), msgAndArgs ...any) bool {
	return asserter.Panics(tb, f, msgAndArgs...)
}

// Empty reports whether object is empty (zero length or nil as applicable).
// It returns true on success.
func Empty(tb TB, object any, msgAndArgs ...any) bool {
	return asserter.Empty(tb, object, msgAndArgs...)
}

// Len reports whether object has the expected length.
// It returns true on success.
func Len(tb TB, object any, length int, msgAndArgs ...any) bool {
	return asserter.Len(tb, object, length, msgAndArgs...)
}
