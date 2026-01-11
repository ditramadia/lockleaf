package model

type Credential struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}
