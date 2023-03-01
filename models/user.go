package models

// UserReq 请求对象
type UserReq struct {
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

// UserInfo 用户对象
type UserInfo struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"comment:用户名;" bson:"username" json:"username" binding:"required"`
	Password string `gorm:"comment:密码;" bson:"password" json:"password" binding:"required"`
	CreateAt MyTime `gorm:"comment:创建时间;type:datetime;" bson:"createAt" json:"createAt"`
	UpdateAt MyTime `gorm:"comment:更新时间;type:datetime;" bson:"updateAt" json:"updateAt"`
}
