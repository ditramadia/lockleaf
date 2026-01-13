package vault

type Field struct {
	Value    string `json:"value"`
	IsSecret bool   `json:"is_secret"`
}

type Credential struct {
	Name   string           `json:"name"`
	Fields map[string]Field `json:"fields"`
}

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
