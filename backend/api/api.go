package api

import (
	authEndpoint "musicboxd/endpoints/auth"

	"github.com/gin-gonic/gin"
)

type Api struct {
	*gin.Engine
}

func NewApi() *Api {
	r := &Api{gin.Default()}

	api := r.RouterGroup.Group("/v1/api")

	auth := api.Group("/auth")
	auth.GET("/login", authEndpoint.LoginEndpoint)
	auth.GET("/callback", authEndpoint.CallbackEndpoint)

	return r
}

func (r *Api) Start() {
	r.Run(":8080")
}
