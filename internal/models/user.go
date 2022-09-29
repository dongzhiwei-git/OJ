package models

import "time"

type User struct {
	UserID     int32     `gorm:"user_id"    json:"user_id"`
	UserName   string    `gorm:"userName"   json:"userName"`
	Email      string    `gorm:"email"      json:"email"`
	Submit     int32     `gorm:"submit"     json:"submit"`
	Solved     int32     `gorm:"solved"     json:"solved"`
	Defunct    string    `gorm:"defunct"    json:"defunct"`
	Ip         string    `gorm:"ip"         json:"ip"`
	Accesstime time.Time `gorm:"accesstime" json:"accesstime"`
	Volume     int32     `gorm:"volume"     json:"volume"`
	Language   int32     `gorm:"language"   json:"language"`
	Password   string    `gorm:"password"   json:"password"`
	RegTime    time.Time `gorm:"regTime"    json:"regTime"`
	Nick       string    `gorm:"nick"       json:"nick"`
	School     string    `gorm:"school"     json:"school"`
	Role       int32     `gorm:"role"       json:"role"`
}

func (User) TableName() string {

	return "user"
}
