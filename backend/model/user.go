package model

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"forum/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var UserDB *db = NewDB()
var UserCounter *utils.Counter = utils.NewCounter()
var Session sessionMap = make(sessionMap)

type User struct {
	id    int
	email string
	name  string
	pswd  string
}

func Register(email, name, pswd string) error {
	if ok := UserDB.Exist(email); ok {
		return errors.New("email already registered")
	}
	newUser := &User{
		id:    UserCounter.Gen(),
		email: email,
		name:  name,
		pswd:  pswd,
	}
	return UserDB.Store(newUser)
}

func LogIn(c *gin.Context) (string, error) {
	obj := struct {
		Email    string `json:"email"`
		Password string `json:"pswd"`
	}{}
	if err := c.BindJSON(&obj); err != nil {
		return "", errors.New("internal error")
	}
	email := obj.Email
	pswd := obj.Password
	if cookie, err := c.Cookie("campus_forum"); err != nil && cookie != "" {
		if _, ok := Session.Get(cookie); ok {
			return cookie, nil
		}
	}
	if ok := UserDB.Exist(email); !ok {
		return "", errors.New("user do not exist")
	}
	truePswd := UserDB.GetPswd(email)
	if truePswd == "" {
		return "", errors.New("internal error")
	}
	if pswd != truePswd {
		return "", errors.New("got wrong password")
	}

	md5 := md5Encode(email + time.Now().String())
	return md5, nil
}

func md5Encode(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
