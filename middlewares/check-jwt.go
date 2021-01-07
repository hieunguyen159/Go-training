package middlewares

import (
	jwt "api/utils"

	"github.com/gin-gonic/gin"
)

func CheckJwt(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.Verify(c)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		next(c)
	}
}
