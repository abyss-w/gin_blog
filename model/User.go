package model

import (
	"encoding/base64"
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2" label:"角色码"` // 角色
}

// 查询用户是否存在
func ContainsUser(name string) int {
	var user User
	DB.Select("id").Where("name = ?", name).Find(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCEED
}

//添加用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err2 := DB.Create(&data).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//钩子
func (u *User) BeforeSave(*gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// 查询用户列表
func GetUsers(pageSize, pageNum int) ([]User, int64) {
	users := make([]User, 0)
	var total int64
	err2 := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil, 0
	}

	return users, total
}

//修改用户
func EditUser(id int, data *User) int {
	var user User
	m := make(map[string]interface{})
	m["name"] = data.Name
	m["role"] = data.Role
	err2 := DB.Model(&user).Where("ID=?", id).Updates(m).Error
	if err2 != nil {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCEED
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err2 := DB.Where("ID=?", id).Delete(&user).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err2 := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err2 != nil {
		log.Fatal(err2)
	}
	encodingPw := base64.StdEncoding.EncodeToString(HashPw)
	return encodingPw

}

// 登录验证
func CheckLogin(name, password string) int {
	var user User
	DB.Where("name=?", name).Find(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCEED
}
