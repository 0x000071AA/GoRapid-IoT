package api

type HTTPResponse struct {
	message string `json:"message"`
	status  int    `json:"status"`
}

func HTTPCreated(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  201,
	}
}

func HTTPGet(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  200,
	}
}

func HTTPBadRequest(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  400,
	}
}

func HTTPDeleted(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  204,
	}
}

func HTTPUpdated(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  200,
	}
}

func HTTPInternalServerError(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  500,
	}
}

func HTTPNotFound(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  404,
	}
}

func HTTPAlreadyExists(msg string) *HTTPResponse {
	return &HTTPResponse{
		message: msg,
		status:  201,
	}
}
