package middleware

import (
	"fmt"
	"musicboxd/hlp"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	fmt.Println(hlp.Envs["FRONTEND_URL"])

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", hlp.Envs["FRONTEND_URL"])
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, Content-Disposition")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
