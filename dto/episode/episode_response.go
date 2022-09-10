package episode

type EpisodeResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" from:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Linkfilm      string `jspon:"thumbnailfilm" gorm:"type:text"`
}
