package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub       int64  `json:"sub"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

type Payload struct {
	Sub       int64  `json:"sub"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	claims := Claims{
		Sub:       data.Sub,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyJwt(secret string, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// package utils

// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// )

// type Header struct {
// 	Alg string `json:"alg"`
// 	Typ string `json:"typ"`
// }

// type Payload struct {
// 	Sub       int64  `json:"sub"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Email     string `json:"email"`
// }

// func CreateJwt(secret string, data Payload) (string, error) {
// 	header := Header{
// 		Alg: "HS256",
// 		Typ: "JWT",
// 	}

// 	byteArrHeader, err := json.Marshal(header)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}

// 	headerBase64 := base64UrlEncode(byteArrHeader)

// 	dataByteArr, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}

// 	payloadBase64 := base64UrlEncode(dataByteArr)

// 	// message := headerBase64 + "." + payloadBase64
// 	message := fmt.Sprintf("%s.%s", headerBase64, payloadBase64)

// 	byteArrSecret := []byte(secret)
// 	byteArrMessage := []byte(message)

// 	h := hmac.New(sha256.New, byteArrSecret)
// 	h.Write(byteArrMessage)

// 	signature := h.Sum(nil)
// 	signatureBase64 := base64UrlEncode(signature)

// 	jwt := fmt.Sprintf("%s.%s.%s", headerBase64, payloadBase64, signatureBase64)
// 	return jwt, nil
// }

// func base64UrlEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }

// Base64

// var s string
// s = "noyan"

// byteArr := []byte(s)
// fmt.Println(byteArr)

// enc := base64.URLEncoding
// enc = enc.WithPadding(base64.NoPadding)
// b64Str := enc.EncodeToString(byteArr)
// decodedStr, err := enc.DecodeString(b64Str)

// if err != nil {
// 	fmt.Println("Error decoding:", err)
// 	return
// }

// fmt.Println(b64Str)
// fmt.Println(decodedStr)

// base64.URLEncoding.WithPadding(base64.NoPadding)

// SHA-256 (Secure Hash Algorithm)

// data := []byte("noyan dey")

// hash := sha256.Sum256(data)

// fmt.Println("Hash after SAH-256", hash)

// HMAC (Hash-based Message Authentication Code)
// HMAC is a cryptographic algorithm that uses a secret key to generate a hash value for a message. It is used to verify the authenticity and integrity of a message.
// HMAC-SHA256

// secret := []byte("my-secret-key")
// message := []byte("hello world")

// h := hmac.New(sha256.New, secret)
// h.Write(message)

// text := h.Sum(nil)
// fmt.Println(text)

// JWT (JSON Web Token) It has three parts: header, payload, and signature.
// JWT is a compact, URL-safe means of representing claims to be transferred between two parties.

// AES
// AES is a symmetric encryption algorithm that uses a secret key to encrypt and decrypt data.
// AES-256

// jwt, err := utils.CreateJwt("my-secret", utils.Payload{
// 	Sub:       88,
// 	FirstName: "Noyan",
// 	LastName:  "Dey",
// 	Email:     "noyandey88@gmail.com",
// })

// if err != nil {
// 	fmt.Println("Error creating JWT:", err)
// 	return
// }

// fmt.Println(jwt)
