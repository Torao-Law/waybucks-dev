package repositories

import (
	"be-waybucks/models"

	"gorm.io/gorm"
)

type TopingRepository interface {
	FindToping() ([]models.Toping, error)
	GetToping(ID int) (models.Toping, error)
	CreateToping(toping models.Toping) (models.Toping, error)
	DeleteToping(toping models.Toping, ID int) (models.Toping, error)
}

func RepositoryToping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindToping() ([]models.Toping, error) {
	var topings []models.Toping
	err := r.db.Find(&topings).Error

	return topings, err
}

func (r *repository) GetToping(ID int) (models.Toping, error) {
	var toping models.Toping
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.First(&toping, ID).Error

	return toping, err
}

func (r *repository) CreateToping(toping models.Toping) (models.Toping, error) {
	err := r.db.Create(&toping).Error

	return toping, err
}

func (r *repository) DeleteToping(toping models.Toping, ID int) (models.Toping, error) {
	err := r.db.Delete(&toping).Error

	return toping, err
}
