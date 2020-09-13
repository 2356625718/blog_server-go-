package controller

import (
	"blog_server/dao"
	"blog_server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ArticlesController struct{

}

func (at *ArticlesController) Router(engine *gin.Engine){
	engine.GET("/articles",at.getArticles)
	engine.GET("/content/:id",at.getContent)
}
//获取所有文章数据
func (at *ArticlesController) getArticles(t *gin.Context){
	var ar model.Articles
	data,err := new(dao.ArticleDao).GetArticles(ar)
	if err != nil{
		log.Fatalf(err.Error())
		return
	}
	t.JSON(http.StatusOK,data)
}
//获取单个文章数据
func (at *ArticlesController) getContent(t *gin.Context){
	data,err := new(dao.ArticleDao).GetContent(t)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	t.JSON(http.StatusOK,data)
}