package comment

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	App_Key = "token"
)

func TokenHandler(id int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": id,
		"exp":  time.Now().Add(time.Hour * time.Duration(10)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(App_Key))
	return tokenString, err
}

func GetToken(tokenString string) (jwt.MapClaims, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return []byte(App_Key), nil })
	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
		return nil, "token错误"
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		fmt.Println("Timing is everything")
		return nil, "token过期"
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return nil, "token错误"
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, ""
	} else {
		fmt.Println(err)
	}
	return nil, ""
}
