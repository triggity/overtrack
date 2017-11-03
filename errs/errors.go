package errs

import (
	"errors"
)

var ErrorNotFound = errors.New("Not Found Error")
var ErrorJsonProcessing = errors.New("JSON error")
