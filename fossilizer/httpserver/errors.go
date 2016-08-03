package httpserver

var (
	ErrInternalServer = ErrHTTP{"internal server error", 500}
	ErrBadRequest     = ErrHTTP{"bad request", 400}
	ErrUnauthorized   = ErrHTTP{"unauthorized", 401}
	ErrNotFound       = ErrHTTP{"not found", 404}
	ErrData           = ErrHTTP{"data required", 400}
	ErrDataLen        = ErrHTTP{"invalid data length", 400}
	ErrCallbackURL    = ErrHTTP{"callback URL required", 400}
)
