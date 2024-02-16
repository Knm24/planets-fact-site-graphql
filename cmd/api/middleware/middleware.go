package middleware

import (
	"github.com/talgat-ruby/planets-fact-site-graphql/cmd/api/types"
	dbT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
)

type middlewareObject struct {
	api types.Api
	db  dbT.DB
}
