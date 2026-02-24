// Package require provides assertion functions for testing.
//
// It mimics the API of the popular github.com/stretchr/testify/require package.
// This package doesn't provide detailed failure messages like testify does, but it provides a simple and consistent API for making assertions in tests.
//
// It uses [*testing.T.Fatal] to report failures, stopping the test execution immediately after a failed assertion.
package require

import "testing"

var _ = (*testing.T)(nil)
