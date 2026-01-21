package status

type status int

const (

	//Informational Responses
	CONTINUE            status = 100
	SWITCHING_PROTOCOLS status = 101
	PROCESSING          status = 102
	EARLY_HINTS         status = 103

	//Successful responses
	OK            status = 200
	CREATED       status = 201
	ACCEPTED      status = 202
	NO_CONTENT    status = 204
	RESET_CONTENT status = 205

	//Redirection message
	MOVED_PERMANENTLY   status = 301
	FOUND               status = 302
	SEE_OTHER           status = 303
	NOT_MODIFIED        status = 304
	PERMANTENT_REDIRECT status = 308

	//Client error

	BAD_REQUEST      status = 400
	UNAUTHORIZED     status = 401
	PAYMENT_REQUIRED status = 402
	FORBIDDEN        status = 403
	NON_FOUND        status = 404

	//???
	IM_A_TEAPOT status = 418
)
