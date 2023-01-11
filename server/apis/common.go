package apis

import (
	"gametools/server/models"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenClaims struct {
	ID       uint64 `json:"id"`
	Username string `json:"name"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

func NewToken(user *models.User) (string, error) {
	now := time.Now()
	expireTime := now.Add(8 * time.Hour)
	claims := TokenClaims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gametools",
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
