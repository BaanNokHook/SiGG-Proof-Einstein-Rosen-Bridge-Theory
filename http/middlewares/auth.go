package middlewares

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/primasio/wormhole/cache"
)

const AuthorizedUserId = "UserId"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		reqToken := c.Request.Header.Get("Authorization")

		if reqToken == "" {
			c.AbortWithStatus(401)
			return
		}

		// Check token validity

		if err, userId := cache.SessionGet(reqToken); err != nil {

			glog.Error("token not exist", err)
			c.AbortWithStatus(500)

		} else {

			if userId == "" {
				c.AbortWithStatus(401)
			} else {

				userIdNum, err := strconv.Atoi(userId)

				if err != nil {
					glog.Error(err)
					c.AbortWithStatus(500)
					return
				}

				c.Set(AuthorizedUserId, uint(userIdNum))

				// User account based access rate limit

				err, reached := rateLimitReached(userId)

				if err != nil {
					glog.Error(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				} else {
					if reached {
						c.AbortWithStatus(http.StatusBadRequest)
					} else {
						c.Next()
					}
				}
			}
		}
	}
}

func rateLimitReached(userId string) (error, bool) {

	cacheType := cache.GetCacheType()

	if cacheType == "memory" {
		return nil, false
	}

	// API access for a single user is limited to 10 times per minute

	currentMinute := int(time.Now().Unix() / 60)

	slotId := "auth_rate_limit_" + userId + "_" + strconv.Itoa(currentMinute)

	cache := cache.GetCache()

	count, err := cache.Increment(slotId, 1)

	if err != nil {
		return err, false
	}

	if count >= 10 {
		return nil, true
	} else {
		cache.Expire(slotId, time.Minute)
	}

	return nil, false
}
