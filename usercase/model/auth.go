package model

import "time"

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gmail    string `json:"gmail"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type Login struct {
	Gmail    string `json:"gmail"`
	Password string `json:"password"`
}

type Reply struct {
	RefreshToken string
	AccessToken  string
}

type JwtConfig struct {
	TokenTimeLife time.Duration
	SecretKet     string
}
