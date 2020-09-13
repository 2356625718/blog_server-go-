package controller

import (
	"blog_server/dao"
	"blog_server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LinkController struct{

}

func (link *LinkController) Router(r *gin.Engine){
	r.GET("/links",link.getLinks)
	r.POST("/addLinks",link.addLinks)
}
//获取友链信息
func (link *LinkController) getLinks(t *gin.Context){
	var ln model.Link
    data,err := new(dao.LinkDao).GetLinks(ln)
	if err != nil{
		log.Fatalf(err.Error())
		return
	}
	t.JSON(http.StatusOK,data)
}
//添加友链
func (link *LinkController) addLinks(t *gin.Context){
	data,_ := new(dao.LinkDao).AddLink(t)
	if data == true{
		t.JSON(http.StatusOK,map[string]interface{}{
			"msg":data,
		})
	}
}
