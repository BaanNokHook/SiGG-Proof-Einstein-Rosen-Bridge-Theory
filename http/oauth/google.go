package oauth

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"context"
	"encoding/json"
	"time"

	"github.com/primasio/wormhole/config"
	"github.com/primasio/wormhole/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig *oauth2.Config

type GoogleUserInfoResponse struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func getGoogleOAuthConfig() *oauth2.Config {
	if googleOAuthConfig != nil {
		return googleOAuthConfig
	}

	c := config.GetConfig()

	clientId := c.GetString("oauth.google.client_id")
	clientSecret := c.GetString("oauth.google.client_secret")

	scheme := c.GetString("application.scheme")
	domain := c.GetString("application.domain")

	googleOAuthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  scheme + "://" + domain + "/v1/oauth/callback/google",
	}

	return googleOAuthConfig
}

func HandleGoogleAuthCallback(code string) (err error, userId uint) {

	googleConfig := getGoogleOAuthConfig()

	// 1. Use code to get Google access token

	ctx, cancelFn1 := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn1()

	token, e := googleConfig.Exchange(ctx, code)

	if e != nil {
		return e, 0
	}

	// 2. Use access token to get user info

	ctx, cancelFn2 := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn2()

	url := "https://www.googleapis.com/oauth2/v2/userinfo"
	client := googleConfig.Client(ctx, token)
	resp, e := client.Get(url)

	if e != nil {
		return e, 0
	}

	defer resp.Body.Close()

	userInfo := &GoogleUserInfoResponse{}

	if e := json.NewDecoder(resp.Body).Decode(userInfo); e != nil {
		return e, 0
	}

	// 3. Process user info

	result := &OAuthResult{
		Type:      models.OAuthGoogle,
		Id:        userInfo.Id,
		Email:     userInfo.Email,
		Name:      userInfo.Name,
		AvatarURL: userInfo.Picture,
	}

	if err, userId := result.Process(); err != nil {
		return err, 0
	} else {
		return nil, userId
	}
}

func HandleGoogleAuth(state string) (redirectUrl string) {

	url := getGoogleOAuthConfig().AuthCodeURL(state)

	return url
}
