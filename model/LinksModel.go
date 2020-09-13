package model

import ()

type Link struct{
	Id int `json"id"`
	Reason string `json:"reason"`
	Url string `json:"url"`
}
