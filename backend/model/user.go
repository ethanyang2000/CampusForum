package model

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"forum/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var UserDB *db = NewDB()
var UserCounter *utils.Counter = utils.NewCounter()
var Session sessionMap = make(sessionMap)

type User struct {
	Id    int    `bson:"id"`
	Email string `bson:"email"`
	Name  string `bson:"name"`
	Pswd  string `bson:"pswd"`
}

func Register(email, name, pswd string) error {
	if ok := UserDB.UserExist(email); ok {
		return errors.New("email already registered")
	}
	newUser := &User{
		Id:    UserCounter.Gen(),
		Email: email,
		Name:  name,
		Pswd:  pswd,
	}
	return UserDB.UserStore(newUser)
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
	if ok := UserDB.UserExist(email); !ok {
		return "", errors.New("user do not exist")
	}
	truePswd, err := UserDB.GetPswd(email)
	if err != nil {
		return "", err
	}
	if truePswd == "" {
		return "", errors.New("internal error")
	}
	if pswd != truePswd {
		return "", errors.New("got wrong password")
	}

	md5 := md5Encode(email + time.Now().String())

	u, err := UserDB.Fetch(email)
	Session.Store(md5, u)
	if err != nil {
		return md5, err
	} else {
		return md5, nil
	}
}

func md5Encode(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func (u *User) ChangeName(newName string) error {
	filter := bson.D{{"email", u.Email}}
	_, err := UserDB.users.UpdateOne(context.Background(), filter, bson.D{
		{"$set", bson.D{{"name", newName}}},
	})
	if err != nil {
		return err
	}
	u.Name = newName
	return nil
}

func (u *User) ChangePswd(newPswd string) error {
	filter := bson.D{{"email", u.Email}}
	_, err := UserDB.users.UpdateOne(context.Background(), filter, bson.D{
		{"$set", bson.D{{"pswd", newPswd}}},
	})
	if err != nil {
		return err
	}
	u.Pswd = newPswd
	return nil
}

func CheckCookie(cookie string) (*User, bool) {
	return Session.Get(cookie)
}

func GetUserByCookie(c *gin.Context) (*User, error) {
	if cookie, err := c.Cookie("campus_forum"); err != nil {
		return nil, err
	} else {
		if user, ok := Session.Get(cookie); !ok {
			return nil, errors.New("load session failed")
		} else {
			return user, nil
		}
	}
}
