package vault

import (
	"testing"
)

// ========================================================================
// AddCredential
// ========================================================================

// TestAddCredential tests the creation of a credential
func TestAddCredential(t *testing.T) {
	// Setup new credential and fields
	name := "Github"
	fields := map[string]Field{
		"username": {Label: "username", Value: "gopher", IsSecret: false},
		"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
	}

	// Create a dummy vault
	v := NewVault("test-vault")

	// Test adding the credential
	err := v.AddCredential(name, fields)
	if err != nil {
		t.Fatalf("Failed to add credential: %v", err)
	}

	// Assert data integrity
	c := v.Credentials["github"]

	if c.Name != "Github" {
		t.Errorf("Missing credential 'github' after adding")
	}

	if c.Fields["username"].Value != "gopher" {
		t.Errorf("Expected value gopher, got %s", c.Fields["username"].Value)
	}
}

// TestAddCredentialWithoutFields tests the creation of a credential without providing the vields
func TestAddCredentialWithoutFields(t *testing.T) {
	// Setup new credential
	name := "Github"

	// Create a dummy vault
	v := NewVault("test-vault")

	// Test adding the credential
	err := v.AddCredential(name, nil)
	if err != nil {
		t.Fatalf("Failed to add credential: %v", err)
	}

	// Assert data integrity
	if c := v.Credentials["github"]; c.Name != "Github" {
		t.Errorf("Missing credential 'github' after adding")
	}

	if c := v.Credentials["github"]; c.Fields != nil {
		t.Errorf("Expected fields to be nil")
	}
}

// ========================================================================
// AddCredential
// ========================================================================

// TestGetField tests the fetching of a field
func TestGetField(t *testing.T) {
	// Setup new credential and fields
	name := "Github"
	fields := map[string]Field{
		"username": {Label: "username", Value: "gopher", IsSecret: false},
		"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
	}

	// Create a dummy vault
	v := NewVault("test-vault")

	// Add a dummy credential
	err := v.AddCredential(name, fields)
	if err != nil {
		t.Fatalf("Failed to add credential: %v", err)
	}

	// Test getting the field
	value, err := v.GetField("Github", "username")
	if err != nil {
		t.Fatalf("Failed to get field: %v", err)
	}

	// Assert data integrity
	if value != "gopher" {
		t.Errorf("Expected value gopher, got %s", value)
	}
}
