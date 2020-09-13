package dao

import (
	"blog_server/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct{
	UserName string `json:"userName"`
	Password string `json:"password"`
}


func (lo *Login) CheckLogin(t *gin.Context){
	err := t.BindJSON(&lo)
	var flag bool = true
	if lo.UserName != "周雨" && lo.Password != "2356625718"{
		flag = false
		t.JSON(http.StatusOK,gin.H{
			"data":flag,
		})
		return
	}
	if err != nil{
		log.Fatalf(err.Error())
		return
	}
	token := lo.Login(t,lo)
	t.JSON(http.StatusOK,gin.H{
		"data":flag,
		"token":token,
	})
}

func (lo *Login) Login(t *gin.Context,lg *Login)(signedToken string){
	claims := &model.JWT{
		Username:lg.UserName,
		Password:lg.Password,

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	Serect := "zhouxiaoyu"
	signedToken,_ = token.SignedString([]byte(Serect))
	return signedToken
}
