package models

import "time"

type Transaction struct {
	ID       int          `json:"id" gorm:"primary_key:auto_increment"`
	Stardate time.Time    `json:"created_at"`
	Duedate  time.Time    `json:"updated_at"`
	Attache  string       `json:"attache" gorm:"type: text"`
	Status   string       `json:"status" gorm:"type: text"`
	User UsersProfileResponse `json:"user"`
	UserID int `json:"-"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname" form:"fullname" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender"  gorm:"type: varchar(255)" `
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: text"`
	Subscribe string `json:"subscribe" gorm:"type:text"`
}

func (UserResponse) TableName() string {
	return "transactions"
}
