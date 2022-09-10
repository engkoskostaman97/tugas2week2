package models

type Film struct {
	ID            int    `json:"id"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string `jspon:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string   `json:"year" gorm:"type:text"`
	
	Description   string `json:"description" gorm:"type: varchar(255)"`
}
