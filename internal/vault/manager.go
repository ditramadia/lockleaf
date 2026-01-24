package vault

import (
	"errors"
	"strings"
)

var (
	ErrCredentialNotFound = errors.New("credential not found")
	ErrCredentialExists   = errors.New("credential already exists")
	ErrFieldNotFound      = errors.New("field not found")
)

// ========================================================================
// Credential Operation
// ========================================================================

// AddCredential handles the logic of adding a credential to the vault
func (v *Vault) AddCredential(name string, fields map[string]Field) error {
	credentialKey := strings.ToLower(name)

	// Create new credentials if it doesn't exist yet
	if v.Credentials == nil {
		v.Credentials = make(map[string]Credential)
	}

	// Validate if credential with the same name already exists
	if _, exists := v.Credentials[credentialKey]; exists {
		return ErrCredentialExists
	}

	// Create Credential data
	v.Credentials[credentialKey] = Credential{
		Name:   name,
		Fields: fields,
	}

	return nil
}

// ========================================================================
// Field Operation
// ========================================================================

// GetField retrieves a specific field value
func (v *Vault) GetField(credName, fieldName string) (string, error) {
	credentialKey := strings.ToLower(credName)
	fieldKey := strings.ToLower(fieldName)

	// Validate if the credential exists
	cred, ok := v.Credentials[credentialKey]
	if !ok {
		return "", ErrCredentialNotFound
	}

	// Validate if the field exists
	field, ok := cred.Fields[fieldKey]
	if !ok {
		return "", ErrFieldNotFound
	}

	return field.Value, nil
}
