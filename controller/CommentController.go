package controller

import (
	"blog_server/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CommentController struct{

}

func (cc *CommentController) Router(g *gin.Engine){
	g.GET("/getComment",cc.getComment)
	g.POST("/commentSubmit/:name/:content",cc.commentSubmit)
	g.GET("/replySubmit",cc.replySubmit)
}

func (cc *CommentController) getComment(t *gin.Context){
	data,_ := new(dao.CommentDao).GetComment()
	t.JSON(http.StatusOK,data)
}

func (cc *CommentController) commentSubmit(t *gin.Context){
	data,err := new(dao.CommentDao).CommentSubmit(t)
	if err != nil{
		log.Fatalf(err.Error())
		return
	}
	t.JSON(http.StatusOK,map[string]interface{}{
		"msg":data,
	})
}

func (cc *CommentController) replySubmit(t *gin.Context){
	data,err := new(dao.CommentDao).ReplySubmit(t)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	t.JSON(http.StatusOK,map[string]interface{}{
		"msg":data,
	})
}
