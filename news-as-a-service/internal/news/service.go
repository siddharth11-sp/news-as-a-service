package news

import "github.com/google/uuid"

type Service interface {
	ListNews(
		entityID uuid.UUID,
		sentiment string,
		source string,
		page int,
		pageSize int,
	) (*ListNewsResponse, error)
}

type service struct {
	repo Repository
}

func NewService(
	repo Repository,
) Service {

	return &service{
		repo: repo,
	}
}

func (s *service) ListNews(
	entityID uuid.UUID,
	sentiment string,
	source string,
	page int,
	pageSize int,
) (*ListNewsResponse, error) {

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	articles,
		total,
		err :=
		s.repo.ListByEntity(
			entityID,
			sentiment,
			source,
			page,
			pageSize,
		)

	if err != nil {
		return nil, err
	}

	return &ListNewsResponse{
		Data:         articles,
		TotalRecords: total,
		Page:         page,
		PageSize:     pageSize,
	}, nil
}
