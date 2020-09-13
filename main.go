package main

import (
	"blog_server/controller"
	"blog_server/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func main(){
	r := gin.Default()
	Cfg,err := tool.ParseConfig("./config/app.json")
	if err != nil{
		log.Fatalf(err.Error())
		return
	}
	r.Use(tool.Cors())
	registerRouter(r)
	r.Run(Cfg.AppHost+":"+Cfg.AppPort)
}

func registerRouter(r *gin.Engine){
	new(controller.ArticlesController).Router(r)
	new(controller.LinkController).Router(r)
	new(controller.CommentController).Router(r)
	new(controller.LoginController).Router(r)
}



