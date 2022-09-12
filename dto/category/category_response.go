package category

type CategoryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" form:"name" validate:"required"`
}
type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}
