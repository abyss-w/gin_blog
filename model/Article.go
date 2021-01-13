package model

import (
	"github.com/abyss-w/gin_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Category Category `gorm:"foreignKey:Cid" json:"category"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

//添加文章
func AddArticle(data *Article) int {
	err2 := DB.Create(&data).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询分类下的所有文章
func GetCateArticle(cid int, pageSize, pageNum int) ([]Article, int, int64) {
	articles := make([]Article, 0)
	var total int64
	err2 := DB.Preload("Category").Where("cid=?", cid).Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articles).Count(&total).Error
	if err2 != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return articles, errmsg.SUCCEED, total
}

//查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err2 := DB.Preload("Category").Where("id=?", id).Find(&article).Error
	if err2 != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}

	return article, errmsg.SUCCEED
}

//查询文章列表
func GetArticles(pageSize, pageNum int) ([]Article, int, int64) {
	articles := make([]Article, 0)
	var total int64
	err2 := DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Count(&total).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}

	return articles, errmsg.SUCCEED, total
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	err2 := DB.Where("id=?", id).Delete(&article).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	m := make(map[string]interface{})
	m["title"] = data.Title
	m["cid"] = data.Cid
	m["desc"] = data.Desc
	m["content"] = data.Content
	m["img"] = data.Img
	err2 := DB.Model(&article).Where("id=?", id).Updates(m).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
