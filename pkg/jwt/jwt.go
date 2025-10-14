package jwt

import (
	"fmt"
	"time"
	"log"

	"github.com/golang-jwt/jwt/v5"
)



func GenerateJWT(id string, updatedAt time.Time,  role string, secret string) (string, error){
	if secret == "" {
		err := fmt.Errorf("secret word jwt is empty")
		log.Printf("[error] %s", err.Error())
		return  "", err
	}
	secretKeyBytes := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": id,
		"updatedAt": updatedAt,
		"role": role,
		"exp": time.Now().Add((time.Hour * 24) * 30).Unix(),
    })

	tokenString, err := token.SignedString(secretKeyBytes)
	if err != nil {
		log.Printf("[error] %s", err.Error())
		return  "", err
	}

	return  tokenString, nil
}