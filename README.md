# testifail
brain dead and limited replacement for https://github.com/stretchr/testify

## How to use this package

### Copy-paste mode

If you don't want to add a dependency, you can:
- simply copy the content of `internal/testifail.go` into your project, and only this file.
- replace the package in the copied file with the package name of your tests (e.g., `package mypackage_test`).
- the file will provide unexported `assert` and `require` variables that you can use in your tests, just like you would with testify.
- You can remove all `stretchr/testify` imports from your project as they are no longer needed, and the local variables will take their place.

### Dependency mode

TBD

## Motivation

[testify](https://github.com/stretchr/testify) is a widely-used testing framework in the Go ecosystem, but it comes with significant complexity. After years of maintenance challenges and debates about API improvements, the historical maintainer has [declared that testify v2 will never happen](https://github.com/stretchr/testify/discussions/1560#discussioncomment-8657735):

> So, as the current only active maintainer, I'm declaring that `v2` will never happen. Or at least a `v2` of the `github.com/stretchr/testify` module with such major breaking changes.

Additionally, the Go team has [discouraged the use of assertion libraries](https://go.dev/doc/faq#assertions), arguing that:

> Go doesn't provide assertions. They are undeniably convenient, but our experience has been that programmers use them as a crutch to avoid thinking about proper error handling and reporting.

However, many codebases have already adopted testify, and migrating away can be time-consuming and risky. Having a simple, drop-in replacement can be valuable for teams seeking to reduce dependencies or simplify their testing approach.

This project was inspired by [a brilliant solution](https://github.com/stretchr/objx/pull/159) from [Emilien Puget](https://github.com/emilien-puget). His approach demonstrated how to replace testify assertions with minimal, locally-defined helpers that preserve the familiar API while removing the external dependency. The elegance of this solution—using simple type declarations to mimic testify's interface—made it clear that a standalone tool could help other projects do the same.

## Alternatives

If you're looking for other ways to move away from testify:

- [go-openapi/testify](https://github.com/go-openapi/testify) - A fork of testify with reduced dependencies and some API improvements

