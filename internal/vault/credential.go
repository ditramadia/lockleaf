package vault

type Credential struct {
	Name   string           `json:"name"`
	Fields map[string]Field `json:"fields"`
}

// NewVault creates an initialized vault struct
func NewCredential(name string, fields map[string]Field) *Credential {
	return &Credential{
		Name:   name,
		Fields: fields,
	}
}
