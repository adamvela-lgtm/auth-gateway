package auth_gateway

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

// GenerateToken generates a new token with the given user ID
func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
		"sub": userID,
	})

	return token.SignedString([]byte("secret"))
}

// ValidateToken validates a given token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err!= nil {
		return nil, err
	}

	return token, nil
}

// GenerateRandomString generates a random string of the given length
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err!= nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// GenerateSession generates a new session
func GenerateSession(r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "session")
	if err!= nil {
		return nil, err
	}

	return session, nil
}

// CloseSession closes a session
func CloseSession(r *http.Request) error {
	session, err := store.Get(r, "session")
	if err!= nil {
		return err
	}

	return session.Save(r, r.URL.Query())
}

func init() {
	var err error
	store, err = sessions.NewCookieStore([]byte("secret"))
	if err!= nil {
		log.Fatal(err)
	}
}