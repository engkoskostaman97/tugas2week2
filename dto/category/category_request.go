package category

type CategoryRequest struct{
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
} 