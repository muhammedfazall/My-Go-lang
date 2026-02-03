package database

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

var Users = map[string]string{}

var Session = map[string]string{}

var Sessionid = "session_id"

func UserExists(u string) bool {
	_, ok := Users[u]
	return ok
}

func CreateUser(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	Users[username] = string(hashed)
	return nil
}

func VerifyUser(username,password string) bool {
	hashed, ok := Users[username]

	if !ok{
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(password)); err != nil{
		return false
	}
	return true
}

func GenerateSessionId(n int) (string,error) {
	b:= make([]byte, n)
	_,err := rand.Read(b)
	if err != nil{
		return "",err
	}
	return hex.EncodeToString(b), nil
}

func CreateSession(username string) (string,error) {
	id,err := GenerateSessionId(16)

	if err != nil{
		return "",err
	}

	Session[id] = username
	return id ,err
}

func DeleteSession(id string)  {
	delete(Session,id)
}