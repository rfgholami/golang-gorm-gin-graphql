package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func Extract(base64EncodedSecret string) (interface{}, interface{}) {

	accessTokenArray := strings.Split(base64EncodedSecret, ".")

	//toString := base64.RawURLEncoding.EncodeToString([]byte(accessTokenArray[2]))
	toString := "ukfFjieO08lY2md-dfxoYr3PFRa6BKj-txPteuL7Cbw"
	//toString := "JLC5XLrIzD1TC64TWnos1c_l3Yw063cDBlU15GrnyqU"

	accessToken := strings.Join([]string{accessTokenArray[0], accessTokenArray[1], string(toString)}, ".")

	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(toString), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		userName := claims["user_name"]
		clientId := claims["client_id"]

		return userName, clientId
	} else {
		fmt.Println("Invalid token claims")
	}

	return nil, nil
}
