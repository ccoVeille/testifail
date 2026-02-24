// Package assert provides assertion functions for testing.
//
// It uses [*testing.T.Error] to report failures, allowing the test to continue execution after a failed assertion.
//
// It mimics the API of the popular github.com/stretchr/testify/assert package.
// This package doesn't provide detailed failure messages like testify does, but it provides a simple and consistent API for making assertions in tests.
//
// Each method returns a boolean indicating whether the assertion succeeded, allowing you to use it in conditional statements if needed.
package assert

import "testing"

var _ = (*testing.T)(nil)
