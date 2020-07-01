package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	hmacSampleSecret = "secret"
	expiredPeriod    = 30
	formatDateLayout = "2006-01-02 15:04:05"
)

func JwtEncoder(userName, customKey string) (string, error) {
	expiredDate := time.Now().Add(time.Second * expiredPeriod)
	claims := jwt.MapClaims{
		"name":      userName,
		"customKey": customKey,
		"expiredAt": expiredDate.Format(formatDateLayout),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

func JwtDecoder(tokenString string) (jwt.MapClaims, error) {
	//Check payload token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSampleSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["name"], claims["customKey"], claims["expiredAt"])
		expiredAt := claims["expiredAt"].(string)

		t, err := time.Parse(formatDateLayout, expiredAt)
		if err != nil {
			return nil, err
		}
		thisTimeString := time.Now().Format(formatDateLayout)
		thisTime, _ := time.Parse(formatDateLayout, thisTimeString)
		diffTime := t.Sub(thisTime).Seconds()

		if diffTime > 0 {
			return claims, nil
		} else {
			return nil, errors.New("Expired Token")
		}

	} else {
		fmt.Println(err)
		return nil, err
	}
}
