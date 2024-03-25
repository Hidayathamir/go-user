package table

import "github.com/sirupsen/logrus"

// User is table `user`. Use this to get table name and column name when query to database.
// Got panic? did you run Init which run initTableUser?
var User *user

type user struct {
	tableName  string
	Dot        *user
	Constraint userConstraint

	ID        string
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
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

	User = &user{
		tableName: "\"user\"",
		Dot:       &user{},
		Constraint: userConstraint{
			UserPk: "user_pk",
			UserUn: "user_un",
		},
		ID:        "id",
		Username:  "username",
		Password:  "password",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}

	User.Dot = &user{
		tableName: User.tableName,
		Dot:       &user{},
		Constraint: userConstraint{
			UserPk: User.Constraint.UserPk,
			UserUn: User.Constraint.UserUn,
		},
		ID:        User.tableName + "." + User.ID,
		Username:  User.tableName + "." + User.Username,
		Password:  User.tableName + "." + User.Password,
		CreatedAt: User.tableName + "." + User.CreatedAt,
		UpdatedAt: User.tableName + "." + User.UpdatedAt,
	}
}
