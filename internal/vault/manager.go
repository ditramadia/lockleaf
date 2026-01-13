package vault

import (
	"errors"
	"fmt"
)

var (
	ErrCredentialNotFound = errors.New("credential not found")
	ErrCredentialExists   = errors.New("credential already exists")
)

// AddCredential handles the logic of adding a credential to the vault
func (v *Vault) AddCredential(name string, fields map[string]Field) error {
	if v.Credentials == nil {
		v.Credentials = make(map[string]Credential)
	}

	// Validate if credential with the same name already exists
	if _, exists := v.Credentials[name]; exists {
		return ErrCredentialExists
	}

	// Create Credential data
	v.Credentials[name] = Credential{
		Name:   name,
		Fields: fields,
	}

	return nil
}

// GetField retrieves a specific field value
func (v *Vault) GetField(credName, fieldName string) (string, error) {
	// Validate if the credential exists
	cred, ok := v.Credentials[credName]
	if !ok {
		return "", ErrCredentialNotFound
	}

	// Validate if the field exists
	field, ok := cred.Fields[fieldName]
	if !ok {
		return "", fmt.Errorf("field '%s' not found", fieldName)
	}

	return field.Value, nil
}
