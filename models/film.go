package models

type Film struct {
	ID            int    `json:"id"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string `jspon:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          int    `json:"year" gorm:"type: int"`
}
