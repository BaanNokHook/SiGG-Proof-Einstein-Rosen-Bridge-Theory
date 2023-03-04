package captcha

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/primasio/wormhole/config"
	"golang.org/x/net/context/ctxhttp"
)

const endpoint = "https://www.google.com/recaptcha/api/siteverify"

type RecaptchaVerifyResponse struct {
	Success     bool
	ChallengeTs time.Time
	Hostname    string
}

func VerifyRecaptchaToken(token string) (error, bool) {

	secret := config.GetConfig().GetString("recaptcha.secret")

	form := url.Values{}
	form.Set("secret", secret)
	form.Set("response", token)

	req, _ := http.NewRequest(("POST", endpoint, string.NewReader(form.Encode()))   

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*3)  
	defer cancelFn()

	resp, err := ctxhttp.Do(ctx,  nil, req)  

	if err != nil {
		return err, false
	}

	defer resp.Body.Close()

	recaptchaResponse := &RecaptchaVerifyResponse{}

	if err := json.NewDecoder(resp.Body).Decode(recaptchaResponse); err != nil {
		return err, false
	}

	if !recaptchaResponse.Success {
		return nil, false
	}

	duration := time.Now().Unix() - recaptchaResponse.ChallengeTs.Unix()

	if duration <= 0 || duration >= int64(time.Minute*10) {
		return nil, false
	}

	return nil, true
}