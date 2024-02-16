package graphql

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/planets-fact-site-graphql/graph/generated"
	"github.com/talgat-ruby/planets-fact-site-graphql/graph/resolver"
)

func (h *Handler) GraphQL(c echo.Context) error {
	ctx := c.Request().Context()
	h.log.InfoContext(ctx, "start Add", "path", c.Path())

	srv := handler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver.Resolver{
					Log: h.log,
					DB:  h.db,
				},
			},
		),
	)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	// srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return nil
}
