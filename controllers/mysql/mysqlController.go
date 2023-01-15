package mysql

import (
	"app/models"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type MySqlController struct{}

// 查找包含title的所有语法
func (con MySqlController) GetGrammarByTitle(c *gin.Context) {
	title := c.Param("title")
	var blue []models.Blues
	var red []models.Reds
	var pre []models.PreExam
	var master []models.Masters
	var dic []models.Dictionaries
	models.DB["grammar"].Where("title LIKE ?", "%"+title+"%").Find(&blue)
	models.DB["grammar"].Where("title LIKE ?", "%"+title+"%").Find(&red)
	models.DB["grammar"].Where("title LIKE ?", "%"+title+"%").Find(&pre)
	models.DB["grammar"].Where("title LIKE ?", "%"+title+"%").Find(&master)
	models.DB["grammar"].Where("title LIKE ?", "%"+title+"%").Find(&dic)
	c.JSON(200, gin.H{
		"blue":   blue,
		"red":    red,
		"pre":    pre,
		"master": master,
		"dic":    dic,
	})
}

// 查找id的漫画
func (con MySqlController) GetComicById(c *gin.Context) {
	id := c.Param("id")
	var comic models.Comics
	models.DB["comic"].Where("id = ?", id).Find(&comic)

	//返回文件夹下所有文件
	files, err := ioutil.ReadDir("./static/files/comic/" + comic.Title + "/")
	if err != nil {
		return
	}
	//填入chapters数组
	var chapters []models.Chapters
	for index, file := range files {
		chapters = append(chapters, models.Chapters{
			ChapterNumber: index,
			Title:         file.Name(),
			SourceUrl:     "/file/comic/" + comic.Title + "/" + file.Name(),
		})
	}

	c.JSON(200, gin.H{
		"comic":    comic,
		"chapters": chapters,
	})
}

// 获取推荐漫画
func (con MySqlController) GetRecommendComic(c *gin.Context) {
	var comic []models.Comics
	models.DB["comic"].Limit(10).Find(&comic)
	c.JSON(200,
		comic,
	)
}

func (con MySqlController) GetComicPages(c *gin.Context) {
	id := c.Param("id")
	chapter := c.Param("chapter")
	var comic models.Comics
	models.DB["comic"].Where("id = ?", id).Find(&comic)

	//返回文件夹下所有文件
	files, err := ioutil.ReadDir("./static/files/comic/" + comic.Title + "/" + chapter + "/")
	if err != nil {
		return
	}
	//填入pages数组
	var pages []string
	for _, file := range files {
		pages = append(pages, "/file/comic/"+comic.Title+"/"+chapter+"/"+file.Name())
	}
	c.JSON(200, gin.H{
		"pages": pages,
	})
}
