package middleware

// ReqContextKey : Key type used to put data in Request Context.
type ReqContextKey string

const (
	// BodyKey : Key for adding request body in request context.
	BodyKey ReqContextKey = "body"
)
