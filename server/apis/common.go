package apis

import (
    "gametools/server/app"
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
    expireTime := app.GetApp().Config.Auth.Token.ExpireTime
    if expireTime <= 0 {
        expireTime = 8 * time.Hour
    }
    expireAt := now.Add(expireTime)
    claims := TokenClaims{
        ID:       user.ID,
        Username: user.Username,
        IsAdmin:  user.IsAdmin,
        RegisteredClaims: jwt.RegisteredClaims{
            Issuer:    "gametools",
            ExpiresAt: jwt.NewNumericDate(expireAt),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(app.GetApp().Config.Auth.Token.SigningKey))
    if err != nil {
        return "", err
    }
    return t, nil
}

func ParseToken() {

}
