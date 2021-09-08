package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Pair struct {
	Key   string
	Value string
}

func GeneralToken(key string, hours int64, pairs []Pair) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(hours)).Unix()
	claims["iat"] = time.Now().Unix()
	for _, pair := range pairs {
		claims[pair.Key] = pair.Value
	}
	token.Claims = claims
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return ss
}

func GeneralForeverCode(key string, pairs []Pair) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	for _, pair := range pairs {
		claims[pair.Key] = pair.Value
	}
	token.Claims = claims
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return ss
}

func DecodeForeverCode(key, token string) (jwt.MapClaims, error) {
	tk, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	claims := tk.Claims.(jwt.MapClaims)
	return claims, err
}

/*
 * 检验token 是否有效
 */
func CheckToken(key, token string) (jwt.MapClaims, bool) {
	tk, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, false
	}
	claims := tk.Claims.(jwt.MapClaims)
	if int64(claims["exp"].(float64)) > time.Now().Unix() {
		return claims, true
	} else {
		return nil, false
	}
}
