package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	ID       uint64 `json:"id"`
	Username string `json:"name"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

func NewToken(cla *TokenClaims) (string, error) {
	now := time.Now()
	expireTime := GetApp().Config.Auth.Token.ExpireTime
	if expireTime <= 0 {
		expireTime = 8 * time.Hour
	}
	expireAt := now.Add(expireTime)
	claims := TokenClaims{
		ID:       cla.ID,
		Username: cla.Username,
		IsAdmin:  cla.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gametools",
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GetApp().Config.Auth.Token.SigningKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseToken(ctx *fiber.Ctx) *TokenClaims {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return &TokenClaims{
		ID:       uint64(claims["id"].(float64)),
		Username: claims["name"].(string),
		Nickname: claims["nickname"].(string),
		IsAdmin:  claims["isAdmin"].(bool),
	}
}
