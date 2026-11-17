package models

type UserModel struct {
	BaseModel
	UserName string `gorm:"size:16" json:"username"`
	NickName string `gorm:"size:32" json:"nickname"`
	Password string `gorm:"size:64" json:"-"`
	RoleID   int8   `json:"roleID"` // 1 管理员 2 普通用户
}
