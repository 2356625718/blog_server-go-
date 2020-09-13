package model

import (
	"database/sql"
)

type Reply struct{
	Id int `json:"id"`
	CommentName sql.NullString `json:"commentName"`
	ReplyId int `json:"replyId"`
	Content string `json:"content"`
}
type Comment struct{
	Id int `json:"id"`
	UserName string `json:"userName"`
	CommentDate string `json:"commentDate"'`
	CommentContent string `json:"commentContent"`
	ReplyArry []Reply `json:"replyArry"`
}
