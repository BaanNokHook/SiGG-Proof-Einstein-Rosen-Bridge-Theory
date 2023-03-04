package api

import (
	"github.com/gin-gonic/gin"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/http/middlewares"
	"github.com/primasio/wormhole/models"
)

type ArticleController struct{}

func (ctrl *ArticleController) Publish(c *gin.Context) {
	var article models.Article

	if err := c.ShouldBind(&article); err != nil {
		Error(err.Error(), c)
	} else {
		dbi := db.GetDb()

		userId, _ := c.Get(middlewares.AuthorizedUserId)

		article.UserID = userId.(uint)

		dbi.Create(&article)

		// TODO: Create async task to publish article to Primas

		Success(article, c)
	}
}

func (ctrl *ArticleController) Get(c *gin.Context) {

}
