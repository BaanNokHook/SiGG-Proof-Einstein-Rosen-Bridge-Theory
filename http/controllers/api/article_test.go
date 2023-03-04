package api_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/tests"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestArticleController_Publish(t *testing.T) {

	PrepareAuthToken(t)

	article, err := tests.CreateTestArticle(systemUser)

	assert.Equal(t, err, nil)

	data := url.Values{}
	data.Set("title", article.Title)
	data.Set("content", article.Content)

	req, _ := http.NewRequest("POST", "/v1/articles", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Authorization", authToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, w.Code, 200)
}
