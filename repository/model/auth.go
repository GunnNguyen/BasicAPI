package model

import "time"

type User struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gmail     string    `json:"gmail"`
	Address   string    `json:"address"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
