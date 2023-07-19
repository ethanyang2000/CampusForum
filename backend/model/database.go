package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
		data: conn.Database("campus_forum").Collection("users"),
	}
}

type db struct {
	data *mongo.Collection
}

func (d *db) Exist(email string) bool {
	user := new(User)
	err := d.data.FindOne(context.Background(), bson.M{"email": email}).Decode(user)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (d *db) Store(u *User) error {
	b := bson.M{
		"email": u.email,
		"name":  u.name,
		"pswd":  u.pswd,
	}
	_, err := d.data.InsertOne(context.Background(), b)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *db) GetPswd(email string) string {
	type u struct {
		Email string `bson:"email"`
		Name  string `bson:"name"`
		Pswd  string `bson:"pswd"`
	}
	var user u

	err := d.data.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		return ""
	} else {
		return user.Pswd
	}
}
