package api

import (
	"musicboxd/api/middleware"
	authEndpoint "musicboxd/endpoints/auth"
	gqlEndpoint "musicboxd/endpoints/graphql"

	"github.com/gin-gonic/gin"
)

type Api struct {
	*gin.Engine
}

func NewApi() *Api {
	r := &Api{gin.Default()}
	r.Use(middleware.CORSMiddleware())

	api := r.RouterGroup.Group("/v1/api")

	auth := api.Group("/auth")
	auth.GET("/login", authEndpoint.LoginEndpoint)
	auth.GET("/callback", authEndpoint.CallbackEndpoint)

	gql := api.Group("/graphql")
	gql.Use(middleware.GinContextToContextMiddleware())
	gql.POST("/query", gqlEndpoint.GraphqlHandler())
	gql.GET("/", gqlEndpoint.PlaygroundHandler())

	return r
}

func (r *Api) Start() {
	r.Run(":8080")
}
