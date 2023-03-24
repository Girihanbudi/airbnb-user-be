package jwt

import (
	"airbnb-user-be/internal/pkg/env"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateToken(creds map[string]interface{}, duration int) *string {
	// Create the JWT claims
	claims := jwt.MapClaims{}

	tokenId, err := gonanoid.New()
	if err != nil {
		return nil
	}

	expirationTime := time.Now().Add(time.Duration(duration))

	claims["jti"] = tokenId                            // token unique id
	claims["iss"] = env.CONFIG.Domain                  // issuer
	claims["exp"] = jwt.NewNumericDate(expirationTime) // expired time

	// added additional claims
	for k, v := range creds {
		if _, ok := claims[k]; ok {
			claims[k] = v
		}
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the generated key using secretKey
	key := []byte(env.CONFIG.Jwt.Secret)
	token, err := jwtToken.SignedString(key)
	if err != nil {
		return nil
	}

	return &token
}

func ExtractTokenMetadata(token string) *jwt.MapClaims {

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unknown signing method")
		}
		key := []byte(env.CONFIG.Jwt.Secret)
		return key, nil
	})
	if err != nil || !jwtToken.Valid {
		return nil
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}
	return &claims
}
