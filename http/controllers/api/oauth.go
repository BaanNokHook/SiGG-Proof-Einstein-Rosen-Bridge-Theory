package api

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/primasio/wormhole/cache"
	"github.com/primasio/wormhole/http/oauth"
	"github.com/primasio/wormhole/http/token"
	"github.com/primasio/wormhole/util"
)

type OAuthController struct{}

func (ctrl *OAuthController) GoogleAuth(c *gin.Context) {

	redirectURI := c.Query("redirect_uri")

	if redirectURI == "" {
		Error("missing query param redirect_uri", c)
		return
	}

	// State is used to prevent attack
	// also as a session key to remember the source of request
	state := util.RandString(8)

	if err := cache.GetCache().Set("oauth_state_"+state, redirectURI, time.Hour*2); err != nil {
		ErrorServer(err, c)
		return
	}

	redirectUrl := oauth.HandleGoogleAuth(state)
	c.Redirect(301, redirectUrl)
}

func (ctrl *OAuthController) GoogleAuthCallback(c *gin.Context) {

	code := c.Query("code")
	state := c.Query("state")
	googleError := c.Query("error")

	var redirectUri string

	// Check state

	if err := cache.GetCache().Get("oauth_state_"+state, &redirectUri); err != nil {

		if err != cache.ErrCacheMiss && err != cache.ErrNotStored {
			ErrorUnauthorized("state expired", c)
		} else {
			ErrorServer(err, c)
		}

		return
	}

	if redirectUri == "" {
		ErrorServer(errors.New("redirect uri not found"), c)
		return
	}

	// Check google return error

	if googleError != "" {
		c.Redirect(301, redirectUri+"?error="+googleError)
		return
	}

	err, userId := oauth.HandleGoogleAuthCallback(code)

	if err != nil {
		ErrorServer(err, c)
		return
	}

	err, accessToken := token.IssueToken(userId, false)

	if err != nil {
		ErrorServer(err, c)
		return
	}

	// Redirect to where it begins

	// TODO: The redirect URL must be pre-registered ones to avoid attack

	c.Redirect(301, redirectUri+"?token="+accessToken.Token)
}
