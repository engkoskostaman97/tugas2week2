package film

type FilmRequest struct {
	ID            int    `json:"id"`
	Title         string `json:"title" from:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string `jspon:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          int    `json:"year" form:"year" gorm:"type: int"`
}
