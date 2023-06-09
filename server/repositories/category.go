package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategory() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	CreateCategory(Category models.Category) (models.Category, error)
	UpdateCategory(Category models.Category) (models.Category, error)
	DeleteCategory(Category models.Category, ID int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategory() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Find(&category).Error

	return category, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category

	err := r.db.First(&category, ID).Error

	return category, err
}

func (r *repository) CreateCategory(Category models.Category) (models.Category, error) {
	err := r.db.Create(&Category).Error

	return Category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	return category, err
}

func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {
	err := r.db.Delete(&category).Error

	return category, err
}
