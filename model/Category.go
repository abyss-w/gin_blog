package model

import (
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID int `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func ContainsCategory(name string) int {
	var category Category
	DB.Select("id").Where("name=?", name).Find(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED
	}
	return errmsg.SUCCEED
}

//添加分类
func AddCategory(data *Category) int {
	err2 := DB.Create(&data).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询分类列表
func GetCategories(pageSize, pageNum int) ([]Category, int64) {
	categories := make([]Category, 0)
	var total int64
	err2 := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Count(&total).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil, 0
	}

	return categories, total
}

//删除分类
func DeleteCategory(id int) int {
	var category Category
	err2 := DB.Where("id=?", id).Delete(&category).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//编辑分类
func EditCategory(id int, data *Category) int {
	var category Category
	err2 := DB.Model(&category).Where("id=?", id).Update("name", data.Name).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//TODO 查询分类下的所有文章
