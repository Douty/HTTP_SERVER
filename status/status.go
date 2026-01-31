package status

import "strconv"

type Status int

const (

	//Informational Responses
	CONTINUE            Status = 100
	SWITCHING_PROTOCOLS Status = 101
	PROCESSING          Status = 102
	EARLY_HINTS         Status = 103

	//Successful responses
	OK            Status = 200
	CREATED       Status = 201
	ACCEPTED      Status = 202
	NO_CONTENT    Status = 204
	RESET_CONTENT Status = 205

	//Redirection message
	MOVED_PERMANENTLY   Status = 301
	FOUND               Status = 302
	SEE_OTHER           Status = 303
	NOT_MODIFIED        Status = 304
	PERMANTENT_REDIRECT Status = 308

	//Client error

	BAD_REQUEST      Status = 400
	UNAUTHORIZED     Status = 401
	PAYMENT_REQUIRED Status = 402
	FORBIDDEN        Status = 403
	NOT_FOUND        Status = 404
	NOT_ALLOWED      Status = 405

	//Server error
	INTERNAL_SERVER_ERROR Status = 500
	NOT_IMPLEMENTED       Status = 501
	BAD_GATEWAY           Status = 502
	SERVICE_UNAVAILABLE   Status = 503

	//???
	IM_A_TEAPOT Status = 418
)

func ToString(status Status) string {
	return strconv.Itoa(int(status))
}
