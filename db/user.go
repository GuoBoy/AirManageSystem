package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	RealName   string `json:"realname,omitempty"`

	IDCard     string `json:"idcard,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	IsSuper bool  `json:"issuper,omitempty" gorm:"default:false"`
}

func (u *User)AddUser()  {
	DB.Create(u)
}

func DelUser(ID string)  {
	var u User
	DB.Where("id=?", ID).Find(&u)
	DB.Delete(&u)
}

func GetAllUser() ([]*User, int64) {
	var users []*User
	tx := DB.Find(&users)
	if tx.Error != nil{
		return nil, 0
	}
	return users, tx.RowsAffected
}

func IsExist(username, password string) (bool, User) {
	var user User
	if err := DB.Where("username=? and password=?", username, password).First(&user).Error;err!=nil{
		return false, user
	}
	return true, user
}
