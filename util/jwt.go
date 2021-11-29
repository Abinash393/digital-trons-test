package util

import (
	"Abinash393/digital-trons-test/model"
	"log"

	"github.com/dgrijalva/jwt-go"
)

const secret = "ultraSecret"

func GenerateToken(user *model.User) string {
	claims := jwt.MapClaims{
		"UID": user.ID,
		"EXP": 15000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}

	return signedToken
}

func DecodeToken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	log.Println(err)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}
