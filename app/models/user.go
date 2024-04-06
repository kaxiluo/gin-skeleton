package models

type User struct {
	ID
	Username string `json:"name"`
	Status   string `json:"status"`
	CreatedAt
}

func (User) TableName() string {
	return "users"
}
