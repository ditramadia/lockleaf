package vault

type Field struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	IsSecret bool   `json:"is_secret"`
}

// NewVault creates an initialized vault struct
func NewField(label, value string, isSecret bool) *Field {
	return &Field{
		Label:    label,
		Value:    value,
		IsSecret: isSecret,
	}
}
