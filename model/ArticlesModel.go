package model

type Articles struct{
	Id int `json:"id"`
	TypeId int `json:"typeId"`
	TypeName string `json:"typeName"`
	Title string `json:"title"`
	Content string `json:"content"`
	Introduce string `json:"introduce"`
	AddYear string `json:"addYear"`
	AddMonth string `json:"addMonth"`
	AddDay string `json:"addDay"`
	ViewCount int `json:"viewCount"`
}
