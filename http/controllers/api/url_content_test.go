package api_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/models"
	"github.com/primasio/wormhole/util"
)

func PrepareURLContent() (error, *models.URLContent) {

	PrepareSystemUser()

	err, domain := PrepareDomain()

	if err != nil {
		return err, nil
	}

	randStr := util.RandString(10)

	urlContent := &models.URLContent{
		URL:    "https://" + domain.Domain + "/12345" + randStr,
		UserID: systemUser.ID,
	}

	urlContent.HashKey = models.GetURLHashKey(urlContent.URL)

	dbi := db.GetDb()
	dbi.Create(&urlContent)

	return nil, urlContent
}

func PrepareURLContentWithUser(user *models.User) (*models.URLContent, error) {

	err, domain := PrepareDomain()

	if err != nil {
		return nil, err
	}

	randStr := util.RandString(10)

	urlContent := &models.URLContent{
		URL:    "https://" + domain.Domain + "/12345" + randStr,
		UserID: user.ID,
	}

	urlContent.HashKey = models.GetURLHashKey(urlContent.URL)

	dbi := db.GetDb()
	dbi.Create(&urlContent)

	return urlContent, nil
}

func TestURLContentController_Get(t *testing.T) {

	err, urlContent := PrepareURLContent()

	assert.Equal(t, err, nil)

	escaped := url.QueryEscape(urlContent.URL)

	log.Println("escaped url: " + escaped)

	req, _ := http.NewRequest("GET", "/v1/urls/url?url="+escaped, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, w.Code, 200)
}

func TestURLContentController_List(t *testing.T) {

	req, _ := http.NewRequest("GET", "/v1/urls", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, w.Code, 200)
}
