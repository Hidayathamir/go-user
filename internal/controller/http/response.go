package http

type baseResponse struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func setResponseBody(data any, err error) baseResponse {
	if err != nil {
		return baseResponse{Data: data, Error: err.Error()}
	}
	return baseResponse{Data: data, Error: nil}
}
