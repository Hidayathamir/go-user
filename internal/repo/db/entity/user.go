package entity

import "time"

// User is entity user, in db it's table `user`.
type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
