package api_test

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/http/server"
	"github.com/primasio/wormhole/models"
	"github.com/primasio/wormhole/tests"
)

var router *gin.Engine
var systemUser *models.User
var authToken string

func TestMain(m *testing.M) {
	before()
	retCode := m.Run()
	os.Exit(retCode)
}

func PrepareSystemUser() {
	if systemUser != nil {
		return
	}

	user, err := tests.CreateTestUser()

	if err != nil {
		log.Fatal(err)
	}

	dbi := db.GetDb()

	if err := user.SetUniqueID(dbi); err != nil {
		log.Fatal(err)
	}

	dbi.Create(&user)

	systemUser = user
}

func PrepareAuthToken(t *testing.T) {

	PrepareSystemUser()

	if authToken != "" {
		return
	}

	w2 := httptest.NewRecorder()

	login := url.Values{}
	login.Set("username", systemUser.Username)
	login.Set("password", "PrimasGoGoGo")
	login.Set("remember", "on")

	req2, _ := http.NewRequest("POST", "/v1/users/auth", strings.NewReader(login.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(login.Encode())))

	router.ServeHTTP(w2, req2)

	responseStr := w2.Body.String()

	log.Println(responseStr)
	assert.Equal(t, w2.Code, 200)

	var returnData map[string]*json.RawMessage

	err := json.Unmarshal([]byte(responseStr), &returnData)
	assert.Equal(t, err, nil)

	var tokenStruct map[string]string

	err = json.Unmarshal(*returnData["data"], &tokenStruct)
	assert.Equal(t, err, nil)

	authToken = tokenStruct["token"]

	log.Println("token: " + authToken)
}

func ResetDB() {

	tables := []string{
		"articles",
		"url_contents",
		"url_content_comments",
		"domain_votes",
		"domains",
	}

	dbi := db.GetDb()
	var sql string

	if db.GetDbType() == db.SQLITE {
		sql = "DELETE FROM"
	} else {
		sql = "TRUNCATE TABLE"
	}

	for _, table := range tables {
		dbi.Exec(sql + " " + table)
	}
}

func before() {
	log.Println("Setting up test environment")
	tests.InitTestEnv("../../../../config/")
	router = server.NewRouter()
}
