package auth

import (
	"time"
	"fmt" 	
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func CreateToken(jwtSecret, userID, role string) (string, error) {
	now := time.Now()
	exp := now.Add(24 * time.Hour)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}

func ParseToken(jwtSecret, tokenStr string) (*Claims, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			// Ensure token uses HMAC signing
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v",
					token.Header["alg"],
				)
			}

			return []byte(jwtSecret), nil
		},
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse token: %w",
			err,
		)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Subject == "" {
		return nil, fmt.Errorf("token missing subject")
	}

	return &claims, nil
}