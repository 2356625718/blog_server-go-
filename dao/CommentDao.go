package dao

import (
	"blog_server/model"
	"blog_server/tool"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	time "time"
)

type CommentDao struct{

}
//获取留言及回复信息
func (cd *CommentDao) GetComment()(string,error){
	db,_ := tool.DbEngine()
	rows,_ := db.Query("select * from comment_reply")
	var reply model.Reply
	var replySli []model.Reply
	var comment model.Comment
	var commentSli []model.Comment
scan:
	if rows.Next(){
		err := rows.Scan(&reply.Id,&reply.CommentName,&reply.ReplyId,&reply.Content)
		if err != nil{
			log.Fatalf(err.Error())
			return "",err
		}
		replySli = append(replySli,reply)
		goto scan
	}
	rows,_= db.Query("select * from comment_to")
scan2:
	if rows.Next(){
		rows.Scan(&comment.Id,&comment.UserName,&comment.CommentDate,&comment.CommentContent)
		commentSli = append(commentSli,comment)
		goto scan2
	}
	for i,com := range commentSli{
		for _,rep := range replySli{
			if rep.ReplyId == com.Id{
				commentSli[i].ReplyArry = append(commentSli[i].ReplyArry,rep)
			}
		}
	}
	dataSli,_ := json.Marshal(commentSli)
	data := string(dataSli)
	return data,nil
}

//添加留言
func (cd *CommentDao) CommentSubmit(t *gin.Context)(bool,error){
	name := t.Param("name")
	content := t.Param("content")
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")
	comment := model.Comment{
		UserName:       name,
		CommentDate:    date,
		CommentContent: content,
		ReplyArry:      nil,
	}
	db,_ := tool.DbEngine()
	_,err := db.Exec("insert into comment_to(user_name,comment_date,comment_content)" +
		"values(?,?,?);",comment.UserName,comment.CommentDate,comment.CommentContent)
	if err != nil{
		log.Fatal(err.Error())
		return false,err
	}
	return true,nil
}

//添加回复
func (cd *CommentDao) ReplySubmit(t *gin.Context)(bool,error){
	db,_ := tool.DbEngine()
	name := t.Query("name")
	content := t.Query("content")
	id := t.Query("id")
	_,err := db.Exec("insert into comment_reply(comment_name,reply_id,content) values(?,?,?)",name,id,content)
	if err != nil{
		log.Fatal(err.Error())
		return false,err
	}
	return true,nil
}