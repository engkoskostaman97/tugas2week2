package models

type Episode struct {
	ID            int                   `json:"id" gorm:"primary_key:auto_increment"`
	Title         string                `json:"title" from:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string                `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Linkfilm      string                `jspon:"linkfilm" gorm:"type:text"`
	User          UsersProfileResponse  `json:"user"`
	UserID        int                   `json:"-"`
	FilmID        int                   `json:"-"`
	Film          FilmsCategoryResponse `json:"film" `
}
