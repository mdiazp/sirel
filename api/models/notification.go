package models

import "time"

type Notification struct {
	Id           int       `json:"id"`
	ToUser       int       `json:"to_user" valid:"Required"`
	Message      string    `json:"message" valid:"Required;MaxSize(1024)"`
	CreationTime time.Time `json:"creation_time" valid:"Required"`
	Readed       bool      `json:"readed"`
}