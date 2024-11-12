package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// ErrorHandlingMiddleware adalah middleware untuk menangani error secara global
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Menangani error secara global
		defer func() {
			if err := recover(); err != nil {
				// Log error
				log.Printf("Recovered from panic: %v", err)

				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Internal server error",
				})

				ctx.Abort()
			}
		}()

		// Melanjutkan ke handler berikutnya
		ctx.Next()

		// Menangani error yang dikembalikan dari handler
		if len(ctx.Errors) > 0 {
			// Log error
			for _, e := range ctx.Errors {
				log.Printf("Error: %v", e.Err)
			}

			// Kirim respons error
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
			})
			ctx.Abort()
		}
	}
}
