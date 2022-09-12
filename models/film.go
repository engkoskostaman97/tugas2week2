package models

type Film struct {
	ID            int                  `json:"id"`
	Title         string               `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string               `jspon:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string               `json:"year" gorm:"type: varchar(255)" `
	Category      CategoryResponse     `json:"category"`
	Description   string               `json:"description" gorm:"type: varchar(255)"`
	User          UsersProfileResponse `json:"user"`
	UserID        int                  `json:"-"`
}
type FilmsCategoryResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"-"`
}

func (FilmsCategoryResponse) TableName() string {
	return "films"
}
