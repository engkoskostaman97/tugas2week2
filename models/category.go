package models

type Category struct {
	ID     int                   `json:"id"`
	Name   string                `json:"name" form:"name" validate:"required"`
	FilmID int                   `json:"film_id"`
	Film   FilmsCategoryResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type CategoryResponse struct {
	ID     int    `gorm:"PrimaryKey"`
	Name   string `json:"name"`
	FilmID int    `json:"-"`
}

func (CategoryResponse) TableName() string {
	return "categories"
}
