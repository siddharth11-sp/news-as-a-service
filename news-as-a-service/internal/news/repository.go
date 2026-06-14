package news

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(article *NewsArticle) error

	ExistsByURLHash(
		hash string,
	) (bool, error)

	ExistsByEntityDedupeKey(
		entityID uuid.UUID,
		key string,
	) (bool, error)
	ListByEntity(
		entityID uuid.UUID,
		sentiment string,
		source string,
		page int,
		pageSize int,
	) ([]NewsArticle, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) Repository {

	return &repository{
		db: db,
	}
}

func (r *repository) ExistsByURLHash(
	hash string,
) (bool, error) {

	var count int64

	err := r.db.
		Model(&NewsArticle{}).
		Where("url_hash = ?", hash).
		Count(&count).
		Error

	return count > 0, err
}

func (r *repository) ExistsByEntityDedupeKey(
	entityID uuid.UUID,
	key string,
) (bool, error) {

	var count int64

	err := r.db.
		Model(&NewsArticle{}).
		Where(
			"entity_id = ? AND dedupe_key = ?",
			entityID,
			key,
		).
		Count(&count).
		Error

	return count > 0, err
}

func (r *repository) Create(
	article *NewsArticle,
) error {

	return r.db.
		Create(article).
		Error
}

func (r *repository) ListByEntity(
	entityID uuid.UUID,
	sentiment string,
	source string,
	page int,
	pageSize int,
) ([]NewsArticle, int64, error) {

	var (
		articles []NewsArticle
		total    int64
	)

	query :=
		r.db.Model(
			&NewsArticle{},
		).
			Where(
				"entity_id = ?",
				entityID,
			)

	if sentiment != "" {

		query =
			query.Where(
				"sentiment = ?",
				sentiment,
			)
	}

	if source != "" {

		query =
			query.Where(
				"source = ?",
				source,
			)
	}

	if err :=
		query.Count(&total).
			Error; err != nil {

		return nil, 0, err
	}

	offset :=
		(page - 1) * pageSize

	err :=
		query.
			Order(
				"published_date DESC",
			).
			Limit(pageSize).
			Offset(offset).
			Find(&articles).
			Error

	return articles, total, err
}
