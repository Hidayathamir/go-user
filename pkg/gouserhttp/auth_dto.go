package gouserhttp

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
