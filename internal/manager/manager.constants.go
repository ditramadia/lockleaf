package manager

import "errors"

var (
	ErrCredentialNotFound = errors.New("credential not found")
	ErrCredentialExists   = errors.New("credential already exists")
	ErrFieldNotFound      = errors.New("field not found")
)
