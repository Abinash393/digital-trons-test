package util

import (
	"Abinash393/digital-trons-test/model"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret = "ultraSecret"

func GenerateToken(user *model.User) string {
	claims := jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}

	return signedToken
}

func DecodeToken(tokenString string) interface{} {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["uid"]
	} else {
		return nil
	}
}
