package controller

import (
	"blog_server/dao"
	"github.com/gin-gonic/gin"
)

type LoginController struct{

}

func (Lg *LoginController) Router(r *gin.Engine){
	r.POST("/login",Lg.checkLogin)
}

func (Lg *LoginController) checkLogin(t *gin.Context){
	new(dao.Login).CheckLogin(t)
}