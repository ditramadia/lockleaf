package vault

type Vault struct {
	Name        string                `json:"name"`
	Credentials map[string]Credential `json:"credentials"`
}

type Credential struct {
	Name   string           `json:"name"`
	Fields map[string]Field `json:"fields"`
}

type Field struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	IsSecret bool   `json:"is_secret"`
}

func NewVault(name string) *Vault {
	return &Vault{
		Name:        name,
		Credentials: make(map[string]Credential),
	}
}

func NewCredential(name string) *Credential {
	return &Credential{
		Name:   name,
		Fields: make(map[string]Field),
	}
}

func NewField(label, value string, isSecret bool) *Field {
	return &Field{
		Label:    label,
		Value:    value,
		IsSecret: isSecret,
	}
}
