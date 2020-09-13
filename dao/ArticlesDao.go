package dao

import (
	"blog_server/model"
	"blog_server/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type ArticleDao struct{
	*tool.DBTool
}
//获取所有文章信息
func (ad *ArticleDao) GetArticles(ar model.Articles)(data string,err error){
	db,err := tool.DbEngine()
	rows,err := db.Query("select id,typeId,typeName,title,content,introduce,addYear,addMonth,addDay,viewCount from articles")
	if err != nil{
		log.Fatalf(err.Error())
		return "",err
	}
	var sum []model.Articles
scan:
	if rows.Next(){
		err := rows.Scan(&ar.Id,&ar.TypeId,&ar.TypeName,&ar.Title,&ar.Content,&ar.Introduce,&ar.AddYear,&ar.AddMonth,&ar.AddDay,&ar.ViewCount)
		if err != nil{
			log.Fatal(err.Error())
			return "",err
		}
		sum = append(sum,ar)
		goto scan
	}
	byte,_ := json.Marshal(sum)
	data = string(byte)
	return
}

//获取单个文章信息,并增加浏览量
func (ad *ArticleDao) GetContent(t *gin.Context)(string,error){
	id := t.Param("id")
	fmt.Println(id)
	db,_ := tool.DbEngine()
	rows,err := db.Query("select * from articles where id=?",id)
	if err != nil{
		log.Fatal(err.Error())
		return "",nil
	}
	var ar model.Articles
	rows.Next()
	err = rows.Scan(&ar.Id,&ar.TypeId,&ar.TypeName,&ar.Title,&ar.Content,&ar.Introduce,&ar.AddYear,&ar.AddMonth,&ar.AddDay,&ar.ViewCount)
	if err !=nil{
		log.Fatal(err.Error())
		return "",nil
	}
	count := ar.ViewCount+1
	_,err = db.Exec("update articles set viewCount=? where id=?",count,id)
	if err !=nil{
		log.Fatal(err.Error())
		return "",nil
	}
	slice,_ := json.Marshal(ar)
	data := string(slice)
	return data,nil
}
