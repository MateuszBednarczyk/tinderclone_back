package services

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"tinderclone_back/src/pkg/domain"
)

type IJwtTokenizer interface {
	GenerateTokens(user domain.User) *Result
	RefreshToken(rawToken string) *Result
	IsTokenValid(rawToken string) *Result
}

type jwtTokenizer struct {
}

type JwtClaims struct {
	Username string      `json:"Username"`
	Role     domain.Role `json:"Role"`
	jwt.StandardClaims
}

func NewJwtTokenizer() *jwtTokenizer {
	return &jwtTokenizer{}
}

func (s *jwtTokenizer) GenerateTokens(user domain.User) *Result {
	baseTokenClaims := JwtClaims{
		user.Username,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	refreshTokenClaims := JwtClaims{
		user.Username,
		user.Role,
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

func (s *jwtTokenizer) RefreshToken(rawToken string) *Result {
	if s.IsTokenValid(rawToken).Content[0] == false {
		return CreateServiceResult("Invalid token", 403, nil)
	}

	claims, err := decodeJwt(rawToken)
	if err != nil {
		return CreateServiceResult("Invalid claims", 403, nil)
	}

	tokens := s.GenerateTokens(domain.User{Role: claims.Role, Username: claims.Username})
	if tokens.Content == nil {
		return CreateServiceResult(tokens.Message, tokens.Code, nil)
	}

	return tokens
}

func (s *jwtTokenizer) IsTokenValid(rawToken string) *Result {
	if strings.TrimSpace(rawToken) == "" {
		return CreateServiceResult("Given token is empty", 403, []interface{}{})
	}

	tokenString := strings.TrimPrefix(rawToken, "Bearer ")
	if tokenString == rawToken {
		return CreateServiceResult("Given token format is invalid", 403, []interface{}{})
	}

	claims, err := decodeJwt(tokenString)
	if err != nil {
		return CreateServiceResult("Invalid token", 403, []interface{}{})
	}

	return CreateServiceResult("Correct token", 200, []interface{}{claims})
}

func decodeJwt(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return token.Claims.(*JwtClaims), nil
}
