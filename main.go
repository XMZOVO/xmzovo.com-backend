package main

import (
	"app/models"
	"app/routers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8081",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	models.ConnectDatabase()
	r.Use(Cors(), gin.Logger(), gin.Recovery(), TlsHandler())
	r.StaticFS("/file", http.Dir("./static/files"))
	routers.MySqlRoutersInit(r)
	// r.Run(":8081")
	r.RunTLS(":8081", "zhongpeiying.com_bundle.pem", "zhongpeiying.com.key")
}
