package constant

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user"}
var TokenKeyCtxKey = &contextKey{"token_key"}
