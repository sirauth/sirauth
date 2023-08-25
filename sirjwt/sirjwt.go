package sirjwt

import "github.com/golang-jwt/jwt/v5"

func CreateToken(secret string, claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString([]byte(secret))
}
