package oauth_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/http/oauth"
	"testing"
)

func TestResponseUnmarshal(t *testing.T) {
	jsonStr := `{
    "id": "12345",
        "email": "test@gmail.com",
        "verified_email": true,
        "name": "chen zhao",
        "given_name": "chen",
        "family_name": "zhao",
        "link": "https://plus.google.com/105681419613076020853",
        "picture": "https://lh5.googleusercontent.com/-YdfbEoumJVE/AAAAAAAAAAI/AAAAAAAAAAo/Iiq2f7lNVFY/photo.jpg",
        "locale": "zh-CN"
}`
	response := oauth.GoogleUserInfoResponse{}

	err := json.Unmarshal([]byte(jsonStr), &response)

	assert.Equal(t, err, nil)

	assert.Equal(t, response.Id, "12345")
	assert.Equal(t, response.Email, "test@gmail.com")
}
