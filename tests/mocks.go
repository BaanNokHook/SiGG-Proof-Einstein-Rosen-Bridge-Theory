package tests

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"log"

	"github.com/primasio/wormhole/models"
	"github.com/primasio/wormhole/util"
)

func CreateTestUser() (*models.User, error) {

	u := &models.User{}

	randStr := util.RandString(5)

	u.Username = "test_user_" + randStr
	u.Nickname = "Test User " + randStr
	u.Password = "PrimasGoGoGo"

	log.Println("Created test user: " + u.Username)

	return u, nil
}

func CreateTestArticle(user *models.User) (*models.Article, error) {

	article := &models.Article{}

	randStr := util.RandString(5)

	article.UserID = user.ID

	article.Title = "Test Article " + randStr
	article.Content = "<p>This is a test article " + randStr + "</p>"

	return article, nil
}
