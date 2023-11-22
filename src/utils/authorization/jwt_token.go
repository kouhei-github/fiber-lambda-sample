package authorization

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateJwtToken(hiddenValue string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"userId": hiddenValue, // ユーザーIDに変更しても良い
		"admin":  true,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Tokenの期限を72時間で設定
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func IsTokenExpired(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// tokenの期限が切れていないか確認
		expireTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expireTime {
			return "", fmt.Errorf("is token expired")
		}
		email := claims["userId"].(string)
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}
