package middleware

import (
	"dainxor/we/base/logger"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jcors "github.com/itsjamie/gin-cors"
)

func CORSMiddleware() gin.HandlerFunc {
	return corsLib()
}

func corsOwn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000, https://fuzzy-fiesta-g6xqxp4w6vw296v-3000.app.github.dev/")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func corsLib() gin.HandlerFunc {
	//front_url := os.Getenv("FRONTEND_URL")
	//proxy_url := os.Getenv("PROXY_URL")

	//allowedOrigins := []string{front_url + ", " + proxy_url + ", https://fuzzy-fiesta-g6xqxp4w6vw296v-3000.app.github.dev"}

	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000/*", "http://127.0.0.1:6969/*", "https://fuzzy-fiesta-g6xqxp4w6vw296v-3000.app.github.dev/*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowWildcard:    true,
		AllowOriginFunc:  func(origin string) bool { logger.Info(origin); return true },
	})
}

func corsJamie() gin.HandlerFunc {
	return jcors.Middleware(jcors.Config{
		Origins:         "http://localhost:3000, https://fuzzy-fiesta-g6xqxp4w6vw296v-3000.app.github.dev/",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Content-Type, Content-Length",
		Credentials:     true,
		ValidateHeaders: false,
	})
}
