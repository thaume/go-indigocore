package httpserver

var (
	ErrInternalServer = ErrHTTP{"internal server error", 500}
	ErrBadRequest     = ErrHTTP{"bad request", 400}
	ErrUnauthorized   = ErrHTTP{"unauthorized", 401}
	ErrNotFound       = ErrHTTP{"not found", 404}
	ErrOffset         = ErrHTTP{"offset must be a positive integer", 400}
	ErrLimit          = ErrHTTP{"limit must be a posive integer", 400}
)
