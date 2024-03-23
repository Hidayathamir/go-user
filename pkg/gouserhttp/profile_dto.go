package gouserhttp

import "time"

// ReqGetProfileByUsername -.
type ReqGetProfileByUsername struct {
	Username string `json:"username"`
}

// ResGetProfileByUsername -.
type ResGetProfileByUsername struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ReqUpdateProfileByUserID -.
type ReqUpdateProfileByUserID struct {
	UserJWT  string `json:"user_jwt"`
	Password string `json:"password"`
}
