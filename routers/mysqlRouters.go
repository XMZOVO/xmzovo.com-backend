package routers

import (
	"app/controllers/mysql"

	"github.com/gin-gonic/gin"
)

func MySqlRoutersInit(r *gin.Engine) {
	grammarGroup := r.Group("/grammar")
	{
		grammarGroup.GET("/title/:title", mysql.MySqlController{}.GetGrammarByTitle)
	}

	comicGroup := r.Group("/comic")
	{
		comicGroup.GET("/id/:id", mysql.MySqlController{}.GetComicById)
		comicGroup.GET("/recommend", mysql.MySqlController{}.GetRecommendComic)
		comicGroup.GET("/pages/:id/:chapter", mysql.MySqlController{}.GetComicPages)
	}
}
