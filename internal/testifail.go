// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright 2025 Emilien Puget <https://github.com/emilien-puget>
// SPDX-FileCopyrightText: Copyright 2026 Christophe Colombier <https://github.com/ccoVeille>

package internal

// You can copy-paste this file in your tests to use the assertion functions.
// When copying, keep the SPDX-License-Identifier and SPDX-FileCopyrightText lines above
// to preserve the required attribution under the MIT license.
//
// Note: the package name must be adapted to the package of your tests.
// Once done, you can mimic the testify packages, by calling assert.Equal(t, expected, actual)
// or require.Nil(t, actual) in your tests.

import (
	"fmt"
	"reflect"
	"testing"
)

// TB is the minimal interface that testifail needs to work, [testing.T] implements it.
//
// this allows anyone to use their own testing framework as long as it implements this interface.
type TB interface {
	Helper()
	Error(args ...any)
	Fatal(args ...any)
}

// interface validations
var (
	_ TB = (testing.TB)(nil)
	_ TB = (*testing.T)(nil)
)

type (
	nonfatal struct{}
	fatal    struct{}
)

var (
	assert  nonfatal
	require fatal
)

// formatMsgArgs formats the msgAndArgs into a single string message.
func formatMsgArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 {
		return ""
	}

	msg, ok := msgAndArgs[0].(string)
	if !ok {
		return fmt.Sprintf("invalid message, first argument must be a string: %#v", msgAndArgs[0])
	}

	if len(msgAndArgs) == 1 {
		return msg
	}

	return fmt.Sprintf(msg, msgAndArgs[1:]...)
}

func (nonfatal) fail(t TB, msgFailure string, msgAndArgs ...any) bool {
	t.Helper()
	if len(msgAndArgs) > 0 {
		msgFailure += ": " + formatMsgArgs(msgAndArgs...)
	}
	t.Error(msgFailure)
	return false
}

func (fatal) fail(t TB, msgFailure string, msgAndArgs ...any) {
	t.Helper()
	if len(msgAndArgs) > 0 {
		msgFailure += ": " + formatMsgArgs(msgAndArgs...)
	}
	t.Fatal(msgFailure)
}

func (a nonfatal) Equal(t TB, expected, actual any, msgAndArgs ...any) bool {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		return a.fail(t, fmt.Sprintf("not equal:\nexpected: %#v\nactual:   %#v", expected, actual), msgAndArgs...)
	}
	return true
}

func (a fatal) Equal(t TB, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		a.fail(t, fmt.Sprintf("not equal:\nexpected: %#v\nactual:   %#v", expected, actual), msgAndArgs...)
	}
}

func (a nonfatal) NotEqual(t TB, expected, actual any, msgAndArgs ...any) bool {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		return a.fail(t, fmt.Sprintf("should not be equal: %#v", actual), msgAndArgs...)
	}
	return true
}

func (a fatal) NotEqual(t TB, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		a.fail(t, fmt.Sprintf("should not be equal: %#v", actual), msgAndArgs...)
	}
}

func isNil(i any) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func (a nonfatal) Nil(t TB, object any, msgAndArgs ...any) bool {
	t.Helper()
	if !isNil(object) {
		return a.fail(t, fmt.Sprintf("expected nil, got: %#v", object), msgAndArgs...)
	}
	return true
}

func (a fatal) Nil(t TB, object any, msgAndArgs ...any) {
	t.Helper()
	if !isNil(object) {
		a.fail(t, fmt.Sprintf("expected nil, got: %#v", object), msgAndArgs...)
	}
}

func (a nonfatal) NotNil(t TB, object any, msgAndArgs ...any) bool {
	t.Helper()
	if isNil(object) {
		return a.fail(t, "expected not nil", msgAndArgs...)
	}
	return true
}

func (a fatal) NotNil(t TB, object any, msgAndArgs ...any) {
	t.Helper()
	if isNil(object) {
		a.fail(t, "expected not nil", msgAndArgs...)
	}
}

func (a nonfatal) True(t TB, value bool, msgAndArgs ...any) bool {
	t.Helper()
	if !value {
		return a.fail(t, "expected true, got false", msgAndArgs...)
	}
	return true
}

func (a fatal) True(t TB, value bool, msgAndArgs ...any) {
	t.Helper()
	if !value {
		a.fail(t, "expected true, got false", msgAndArgs...)
	}
}

func (a nonfatal) False(t TB, value bool, msgAndArgs ...any) bool {
	t.Helper()
	if value {
		return a.fail(t, "expected false, got true", msgAndArgs...)
	}
	return true
}

func (a fatal) False(t TB, value bool, msgAndArgs ...any) {
	t.Helper()
	if value {
		a.fail(t, "expected false, got true", msgAndArgs...)
	}
}

func (a nonfatal) NoError(t TB, err error, msgAndArgs ...any) bool {
	t.Helper()
	if err != nil {
		return a.fail(t, fmt.Sprintf("expected no error, got: %v", err), msgAndArgs...)
	}
	return true
}

func (a fatal) NoError(t TB, err error, msgAndArgs ...any) {
	t.Helper()
	if err != nil {
		a.fail(t, fmt.Sprintf("expected no error, got: %v", err), msgAndArgs...)
	}
}

func (a nonfatal) Error(t TB, err error, msgAndArgs ...any) bool {
	t.Helper()
	if err == nil {
		return a.fail(t, "expected error, got nil", msgAndArgs...)
	}
	return true
}

func (a fatal) Error(t TB, err error, msgAndArgs ...any) {
	t.Helper()
	if err == nil {
		a.fail(t, "expected error, got nil", msgAndArgs...)
	}
}

func (a nonfatal) Panics(t TB, f func(), msgAndArgs ...any) bool {
	t.Helper()

	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r == nil {
				_ = a.fail(t, "expected panic, but function did not panic", msgAndArgs...)
				return
			}
			didPanic = true
		}()
		f()
	}()
	return didPanic
}

func (a fatal) Panics(t TB, f func(), msgAndArgs ...any) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			a.fail(t, "expected panic, but function did not panic", msgAndArgs...)
		}
	}()
	f()
}

func (a nonfatal) Empty(t TB, object any, msgAndArgs ...any) bool {
	t.Helper()
	if !isEmpty(object) {
		return a.fail(t, fmt.Sprintf("expected empty, got: %#v", object), msgAndArgs...)
	}
	return true
}

func (a fatal) Empty(t TB, object any, msgAndArgs ...any) {
	t.Helper()
	if !isEmpty(object) {
		a.fail(t, fmt.Sprintf("expected empty, got: %#v", object), msgAndArgs...)
	}
}

func isEmpty(object any) bool {
	if object == nil {
		return true
	}
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return true
		}
		return isEmpty(v.Elem().Interface())
	}
	// numbers and structs are never considered empty here
	return false
}

func (a nonfatal) Len(t TB, object any, length int, msgAndArgs ...any) bool {
	t.Helper()
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if v.Len() != length {
			return a.fail(t, fmt.Sprintf("unexpected length, expected %d got %d", length, v.Len()), msgAndArgs...)
		}
		return true
	default:
		return a.fail(t, fmt.Sprintf("Len not supported for kind %s", v.Kind()), msgAndArgs...)
	}
}

func (a fatal) Len(t TB, object any, length int, msgAndArgs ...any) {
	t.Helper()
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if v.Len() != length {
			a.fail(t, fmt.Sprintf("unexpected length, expected %d got %d", length, v.Len()), msgAndArgs...)
		}
	default:
		a.fail(t, fmt.Sprintf("Len not supported for kind %s", v.Kind()), msgAndArgs...)
	}
}
