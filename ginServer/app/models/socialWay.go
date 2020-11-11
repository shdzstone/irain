package models

type SocialWay struct {
	Model
	AdminId uint	`json:"admin_id,omitempty"`
	Name string		`json:"name,omitempty"`
	Url string		`json:"url,omitempty"`
}