package model

type User struct {
	UserID   string `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"notnull unique" binding:"required"`
	Password string `json:"password" gorm:"notnull" binding:"required"`
}
