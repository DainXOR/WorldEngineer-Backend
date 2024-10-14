package middleware

import (
	ccors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jcors "github.com/itsjamie/gin-cors"
)

func CORSMiddleware() gin.HandlerFunc {
	return corsJamie()
}

func corsOwn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000, https://fuzzy-fiesta-g6xqxp4w6vw296v-3000.app.github.dev/")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func corsLib() gin.HandlerFunc {
	return ccors.New(ccors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}

func corsJamie() gin.HandlerFunc {
	return jcors.Middleware(jcors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Content-Type, Content-Length",
		Credentials:     true,
		ValidateHeaders: false,
	})
}
