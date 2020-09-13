package dao

import (
	"blog_server/model"
	"blog_server/tool"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type LinkDao struct{

}
//获取友链数据
func (ln *LinkDao) GetLinks(md model.Link)(string,error){
	db,err := tool.DbEngine()
	if err != nil{
		log.Fatalf(err.Error())
		return "",err
	}
	rows,err := db.Query("select id,reason,url from links")
	if err != nil{
		log.Fatal(err.Error())
		return "", err
	}
	var sum []model.Link
scan:
	if rows.Next(){
		err := rows.Scan(&md.Id,&md.Reason,&md.Url)
		if err != nil{
			return "",err
		}
		sum = append(sum,md)
		goto scan
	}
	slice,_ := json.Marshal(sum)
	data := string(slice)
	return data,nil
}

//增加友链数据
func (ln *LinkDao) AddLink(t *gin.Context)(bool,error){
	reason := t.Query("reason")
	url := t.Query("url")
	db,err := tool.DbEngine()
	if err != nil{
		log.Fatalf(err.Error())
		return false,err
	}
	_,err = db.Exec("insert into links(reason,url)"+
		"values(?,?);",reason,url)
	if err != nil{
		log.Fatal(err.Error())
		return false,err
	}
	return true,nil
}
