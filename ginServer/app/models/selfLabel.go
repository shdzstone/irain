package models

type SelfLabel struct {
	Model
	AdminId uint	`json:"admin_id,omitempty"`
	Name string		`json:"name,omitempty"`
}