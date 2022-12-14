package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisodes() ([]models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	CreateEpisode(Episode models.Episode) (models.Episode, error)
	UpdateEpisode(Episode models.Episode) (models.Episode, error)
	DeleteEpisode(Episode models.Episode) (models.Episode, error)
}

func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpisodes() ([]models.Episode, error) {
	var episodes []models.Episode
	err := r.db.Find(&episodes).Error // add this code

	return episodes, err
}

func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var episode models.Episode	
	err := r.db.First(&episode, ID).Error // add this code

	return episode, err
}
func (r *repository) CreateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Create(&episode).Error

	return episode, err
}

func (r *repository) UpdateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Save(&episode).Error

	return episode, err
}

func (r *repository) DeleteEpisode(episode models.Episode) (models.Episode, error) {

	err := r.db.Delete(&episode).Error

	return episode, err
}
