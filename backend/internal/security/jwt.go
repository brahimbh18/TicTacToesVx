package security

import (
"time"

"github.com/golang-jwt/jwt/v5"
"github.com/google/uuid"
)

type JWTService struct {
secret []byte
}

func NewJWTService(secret string) *JWTService {
return &JWTService{secret: []byte(secret)}
}

func (s *JWTService) CreateToken(userID uuid.UUID) (string, error) {
claims := jwt.MapClaims{
"sub": userID.String(),
"exp": time.Now().Add(24 * time.Hour).Unix(),
"iat": time.Now().Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString(s.secret)
}

func (s *JWTService) ParseToken(tokenStr string) (uuid.UUID, error) {
tok, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
return s.secret, nil
})
if err != nil || !tok.Valid {
return uuid.Nil, err
}
claims, ok := tok.Claims.(jwt.MapClaims)
if !ok {
return uuid.Nil, jwt.ErrTokenInvalidClaims
}
sub, ok := claims["sub"].(string)
if !ok {
return uuid.Nil, jwt.ErrTokenInvalidClaims
}
return uuid.Parse(sub)
}
