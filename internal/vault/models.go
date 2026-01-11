package vault

type Field struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	IsSecret bool   `json:"is_secret"`
}

type Credential struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Vault struct {
	Name        string       `json:"name"`
	Credentials []Credential `json:"credentials"`
}

// NewVault creates an initialized vault struct
func NewVault(name string) *Vault {
	return &Vault{
		Name:        name,
		Credentials: make([]Credential, 0),
	}
}
