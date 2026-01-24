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
