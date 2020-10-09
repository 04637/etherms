package model

type User struct {
	UserCode string `gorm:"column:user_code;type:varchar(10);primaryKey"`
	AuthorId string `gorm:"column:author_id;type:varchar(20);uniqueIndex"`
	RealName string `gorm:"column:real_name;type:varchar(10)"`
}
