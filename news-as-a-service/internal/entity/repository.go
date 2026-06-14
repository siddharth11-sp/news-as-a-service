package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateEntity(entity *Entity) error
	GetEntity(id uuid.UUID) (*Entity, error)
	ListEntities() ([]Entity, error)
	UpdateEntity(entity *Entity) error
	UpdateLastIngestedAt(id uuid.UUID, t time.Time) error
	ListActiveEntities() ([]Entity, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateEntity(entity *Entity) error {
	return r.db.Create(entity).Error
}

func (r *repository) GetEntity(id uuid.UUID) (*Entity, error) {

	var entity Entity

	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *repository) ListEntities() ([]Entity, error) {

	var entities []Entity

	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *repository) UpdateEntity(entity *Entity) error {
	return r.db.Save(entity).Error
}

func (r *repository) UpdateLastIngestedAt(
	id uuid.UUID,
	t time.Time,
) error {

	return r.db.
		Model(&Entity{}).
		Where("id = ?", id).
		Update(
			"last_ingested_at",
			t,
		).Error
}

func (r *repository) ListActiveEntities() ([]Entity, error) {

	var entities []Entity

	err := r.db.
		Where("status = ?", "ACTIVE").
		Find(&entities).
		Error

	return entities, err
}
