package status

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
	NON_FOUND        Status = 404
	NOT_ALLOWED      Status = 405

	//???
	IM_A_TEAPOT Status = 418
)
