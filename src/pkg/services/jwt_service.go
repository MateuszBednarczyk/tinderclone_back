package services

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IJwtTokenizer interface {
	GenerateTokens(username string, isAdmin bool) *Result
	RefreshToken(rawToken string, isAdmin bool) *Result
	IsTokenValid(rawToken string) *Result
}

type jwtTokenizer struct {
}

type JwtClaims struct {
	Username string `json:"name"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func NewJwtTokenizer() *jwtTokenizer {
	return &jwtTokenizer{}
}

func (s *jwtTokenizer) GenerateTokens(username string, isAdmin bool) *Result {
	baseTokenClaims := JwtClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	refreshTokenClaims := JwtClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, baseTokenClaims)
	token, err := rawToken.SignedString([]byte("secret"))
	if err != nil {
		return CreateServiceResult("Invalid token", 500, []interface{}{})
	}

	rawRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err := rawRefreshToken.SignedString([]byte("secret"))
	if err != nil {
		return CreateServiceResult("Invalid token", 403, nil)

	}

	content := []interface{}{map[string]string{"token": "Bearer " + token, "refresh": "Bearer " + refreshToken}}

	return CreateServiceResult("Tokens generated successfully", 200, content)
}

func (s *jwtTokenizer) RefreshToken(rawToken string, isAdmin bool) *Result {
	if s.IsTokenValid(rawToken).Content[0] == false {
		return CreateServiceResult("Invalid token", 403, nil)
	}

	tokens := s.GenerateTokens(rawToken, isAdmin)
	if tokens.Content == nil {
		return CreateServiceResult(tokens.Message, tokens.Code, nil)
	}

	return tokens
}

func (s *jwtTokenizer) IsTokenValid(rawToken string) *Result {
	var isTokenCorrect bool

	if rawToken == "" {
		isTokenCorrect = false
		return CreateServiceResult("Given token is empty", 403, []interface{}{isTokenCorrect})
	}

	tokenString := strings.TrimPrefix(rawToken, "Bearer ")
	if tokenString == rawToken {
		return CreateServiceResult("Given token format is invalid", 403, []interface{}{})
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return CreateServiceResult("Invalid token", 403, []interface{}{})
	}
	if !token.Valid {
		return CreateServiceResult("Invalid token", 403, []interface{}{})
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return CreateServiceResult("Invalid token", 403, []interface{}{})
	}
	isTokenCorrect = true

	return CreateServiceResult("Correct token", 200, []interface{}{isTokenCorrect, claims})
}
