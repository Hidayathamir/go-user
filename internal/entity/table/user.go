package table

import "github.com/sirupsen/logrus"

// User is table `user`. Use this to get table name and column name when query to database.
var User *user

type user struct {
	tableName string

	ID        string
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func (u *user) String() string {
	return u.tableName
}

func initTableUser() {
	if User != nil {
		logrus.Warn("table User already initialized")
		return
	}

	// We need to use \"user\" because user is a reserved keyword.
	//
	// Error:   select user.id from user
	// Success: select "user".id from "user"
	User = &user{
		tableName: "\"user\"",
		ID:        "\"user\".id",
		Username:  "\"user\".username",
		Password:  "\"user\".password",
		CreatedAt: "\"user\".created_at",
		UpdatedAt: "\"user\".updated_at",
	}
}
