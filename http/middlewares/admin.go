package middlewares

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"github.com/gin-gonic/gin"
	"github.com/primasio/wormhole/config"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		reqToken := c.Request.Header.Get("Authorization")

		if reqToken == "" {
			c.AbortWithStatus(401)
			return
		}

		// Check admin token validity

		adminToken := config.GetConfig().GetString("admin.key")

		if adminToken == "" || reqToken != adminToken {
			c.AbortWithStatus(401)
		} else {
			c.Next()
		}
	}
}
