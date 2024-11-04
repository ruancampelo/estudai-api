package model

type Content struct {
	Id          int64  `gorm:"column:id;primarykey;autoIncrement"`
	Title       string `gorm:"column:title;"`
	JsonContent string `gorm:"column:content"`
	//UserId      int64  `gorm:"column:user_id;"`
}
