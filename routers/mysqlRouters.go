package routers

import (
	"app/controllers/mysql"

	"github.com/gin-gonic/gin"
)

func MySqlRoutersInit(r *gin.Engine) {
	grammarGroup := r.Group("/grammar")
	{
		grammarGroup.GET("/title/:title", mysql.MySqlController{}.GetGrammarByTitle)
		grammarGroup.POST("/feedback", mysql.MySqlController{}.PostFeedback)
		grammarGroup.GET("feedback", mysql.MySqlController{}.GetFeedback)
		grammarGroup.POST("/feedback/like/:id", mysql.MySqlController{}.LikeFeedback)
	}

	comicGroup := r.Group("/comic")
	{
		comicGroup.GET("/id/:id", mysql.MySqlController{}.GetComicById)
		comicGroup.GET("/recommend", mysql.MySqlController{}.GetRecommendComic)
		comicGroup.GET("/pages/:id/:chapter", mysql.MySqlController{}.GetComicPages)
	}
}
