package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identify      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogoutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
