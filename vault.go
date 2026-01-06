package main

const (
	FieldEmail    FieldType = "email"
	FieldPhone    FieldType = "phone"
	FieldUsername FieldType = "username"
	FieldPassword FieldType = "password"
	FieldPIN      FieldType = "pin"
)

type FieldType string

type Field struct {
	Key   string    `json:"key"`
	Type  FieldType `json:"type"`
	Value string    `json:"value"`
}

type Credential struct {
	ID     string  `json:"id"`
	Label  string  `json:"label"`
	Fields []Field `json:"fields"`
}

type Vault struct {
	Entries []Credential `json:"entries"`
}
