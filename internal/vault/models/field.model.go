package model

type Field struct {
	Name     string `json:"name"`
	IsSecret bool   `json:"is_secret"`
}
