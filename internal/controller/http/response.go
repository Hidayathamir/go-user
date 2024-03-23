package http

// ResError -.
type ResError struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

// ResString -.
type ResString struct {
	Data  string `json:"data"`
	Error any    `json:"error"`
}
