package router

import (
	"encoding/json"
	"httpserver/request"
	"httpserver/status"
	"slices"
)

func APIGetAllUsers(ctx Context) ([]byte, *HTTPError) {
	allowedMethods := []request.Method{request.GET}

	if !slices.Contains(allowedMethods, ctx.Method) {
		return nil, &HTTPError{Message: "Method Not Allowed", StatusCode: status.NOT_ALLOWED}
	}
	users := []string{"Alice", "Bob", "Charlie"}
	data, err := json.Marshal(users)
	if err != nil {
		return nil, &HTTPError{
			Message:    "Internal Server Error",
			StatusCode: status.INTERNAL_SERVER_ERROR,
		}
	}
	return data, nil
}
