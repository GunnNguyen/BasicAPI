package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Gmail string `json:"gmail"`
	jwt.StandardClaims
}

func GenerateJwt(ctx context.Context, SecretKet string, TokenTimeLife time.Duration, name string, gmail string) (string, error) {
	if SecretKet == "" {
		return "", fmt.Errorf("Secret key not found")
	}
	claims := &JwtCustomClaims{
		Name:  name,
		Gmail: gmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * TokenTimeLife).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SecretKet))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
