package vault

type Vault struct {
	Name        string                `json:"name"`
	Credentials map[string]Credential `json:"credentials"`
}

// NewVault creates an initialized vault struct
func NewVault(name string) *Vault {
	return &Vault{
		Name:        name,
		Credentials: make(map[string]Credential),
	}
}
