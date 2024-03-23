package gouserhttp

import "time"

// ReqLoginUser -.
type ReqLoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ResLoginUser -.
type ResLoginUser struct {
	UserJWT string `json:"user_jwt"`
}

// ReqRegisterUser -.
type ReqRegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ResRegisterUser -.
type ResRegisterUser struct {
	UserID int64 `json:"user_id"`
}

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
