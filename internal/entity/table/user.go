package table

import "github.com/sirupsen/logrus"

// User is table `user`. Use this to get table name and column name when query to database.
// Got panic? did you run Init which run initTableUser?
var User *user

type user struct {
	tableName string

	ID        string
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string

	Constraint userConstraint
}

type userConstraint struct {
	UserPk string
	UserUn string
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
		ID:        "id",
		Username:  "username",
		Password:  "password",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		Constraint: userConstraint{
			UserPk: "user_pk",
			UserUn: "user_un",
		},
	}
}
