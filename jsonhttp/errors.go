package jsonhttp

var (
	ErrInternalServer = ErrHTTP{"internal server error", 500}
	ErrBadRequest     = ErrHTTP{"bad request", 400}
	ErrUnauthorized   = ErrHTTP{"unauthorized", 401}
	ErrNotFound       = ErrHTTP{"not found", 404}
)
