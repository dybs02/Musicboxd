package graphql

import (
	"context"
	"fmt"
	"musicboxd/graph"
	resolver "musicboxd/graph/resolvers"
	"musicboxd/hlp"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == r.Header.Get("Host") {
				return true
			}

			// TODO Remove, after testing in production
			fmt.Printf("Origin: %s, Host: %s\n", origin, r.Header.Get("Host"))
			fmt.Printf("Allowed Origins: %v\n", hlp.Envs["FRONTEND_URL"])
			return slices.Contains([]string{hlp.Envs["FRONTEND_URL"]}, origin)
		},
	}

	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader:              upgrader,
	})

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	// TODO Disable in production
	h.Use(extension.Introspection{})
	h.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		fmt.Printf("\n\n\nOperation: %s \n%s \n%s", oc.OperationName, oc.RawQuery, oc.Variables)
		return next(ctx)
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/v1/api/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
