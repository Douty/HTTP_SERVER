package router

import (
	"encoding/json"
	"httpserver/request"
	"httpserver/status"
	"slices"
)

func APIGetAllUsers(ctx Context) (Asset, *HTTPError) {
	allowedMethods := []request.Method{request.GET}

	if !slices.Contains(allowedMethods, ctx.Method) {
		return Asset{}, &HTTPError{Message: "Method Not Allowed", StatusCode: status.NOT_ALLOWED}
	}
	users := []string{"Alice", "Bob", "Charlie"}
	data, err := json.Marshal(users)
	if err != nil {
		return Asset{}, &HTTPError{
			Message:    "Internal Server Error",
			StatusCode: status.INTERNAL_SERVER_ERROR,
		}
	}
	return Asset{Content: data, ContentType: "application/json"}, nil
}
