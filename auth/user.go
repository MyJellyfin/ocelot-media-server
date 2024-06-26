package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"ocelot/config"
	"ocelot/storage"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	SessionToken   string `json:"session_token"`
	ExpiresAt      int64  `json:"expires_at"`
	JwtTokenString string `json:"jwt_token_string"`
	Type           int    `json:"type"`
}

func (u *User) Login(username string, password string, device string, deviceName string, clientName string, clientVersion string) error {
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}

	h := sha256.New()

	h.Write([]byte(password))
	hashedPass := base64.URLEncoding.EncodeToString(h.Sum(nil))

	user, err := queries.GetUserWithPassword(context.Background(), storage.GetUserWithPasswordParams{
		Username: username,
		Password: string(hashedPass),
	})

	if err != nil {
		return errors.New("No user found.")
	}

	session, err := createSession(int(user.ID), deviceName, clientName, clientVersion, device)

	claims := &jwt.MapClaims{
		"exp": session.ExpiresAt.Unix(),
		"aud": "http://localhost:8080",
		"iss": "ocelotsoftware.org",
		"data": map[string]string{
			"id":       fmt.Sprint(user.ID),
			"username": user.Username,
			"session":  fmt.Sprint(session.ID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var config config.Config
	config.Read()
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return err
	}
	u.JwtTokenString = tokenString

	return nil
}
