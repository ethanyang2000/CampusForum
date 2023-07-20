package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type db struct {
	users *mongo.Collection
	posts *mongo.Collection
}

func NewDB() *db {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := mongo.Connect(ctx,
		options.Client().
			// address
			ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	err = conn.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return &db{
		users: conn.Database("campus_forum").Collection("users"),
		posts: conn.Database("campus_forum").Collection("posts"),
	}
}

func (d *db) UserExist(email string) bool {
	user := new(User)
	err := d.users.FindOne(context.Background(), bson.M{"email": email}).Decode(user)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (d *db) UserStore(u *User) error {
	_, err := d.users.InsertOne(context.Background(), u)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *db) GetPswd(email string) (string, error) {
	user := new(User)
	err := d.users.FindOne(context.Background(), bson.D{{"email", email}}).Decode(user)

	if err != nil {
		return "", err
	} else {
		return user.Pswd, nil
	}
}

func (d *db) Fetch(email string) (*User, error) {
	user := new(User)
	err := d.users.FindOne(context.Background(), bson.D{{"email", email}}).Decode(user)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
