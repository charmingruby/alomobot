package cmd

import "errors"

var (
	ErrNonexistentKind = errors.New("kind does not exists")
	ErrInvalidParam    = errors.New("invalid param")
)
