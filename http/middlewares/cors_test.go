package middlewares_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/http/server"
	"github.com/primasio/wormhole/tests"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestCors(t *testing.T) {
	tests.InitTestEnv("../../config/")
	router := server.NewRouter()

	w := httptest.NewRecorder()

	form := url.Values{}
	form.Set("url", "http://test.com/test.html")
	form.Set("content", "<p>The comment of a URL.</p>")

	req, _ := http.NewRequest("POST", "/v1/comments", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("Origin", "http://test.com")

	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, w.Code, 401)
}
