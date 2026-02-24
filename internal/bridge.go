package internal

// this file is used to bridge the internal package with the public assert and require packages

var (
	Assert  = assert
	Require = require
)
