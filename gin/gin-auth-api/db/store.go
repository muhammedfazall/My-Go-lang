package db

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

var Users = map[string]string{}

var Sessions = map[string]string{}

const SessionCookieName = "sessions_id"

func CreateUser(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	Users[username] = string(hashed)
	return nil
}

func UserExists(username string) bool {
	_, ok := Users[username]
	return ok
}

func VerifyUser(username, password string) bool {
	hashed, ok := Users[username]
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}

func GenerateSessionId(nBytes int) (string, error) {
	b := make([]byte, nBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func CreateSession(username string) (string, error) {
	id, err := GenerateSessionId(16)
	if err != nil {
		return "", err
	}

	Sessions[id] = username
	return id, nil
}

func GetUsernameBySession(id string) (string, bool) {
	username, ok := Sessions[id]
	return username, ok
}

func DeleteSession(id string) {
	delete(Sessions, id)
}
