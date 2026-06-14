package entity

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Service interface {
	CreateEntity(req CreateEntityRequest) (*Entity, error)
	GetEntity(id string) (*Entity, error)
	ListEntities() ([]Entity, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateEntity(
	req CreateEntityRequest,
) (*Entity, error) {

	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if req.IngestionIntervalMinutes <= 0 {
		return nil, errors.New(
			"ingestion interval must be greater than zero",
		)
	}

	aliases := normalizeAliases(req.Aliases)

	entity := &Entity{
		ID:                       uuid.New(),
		Name:                     strings.TrimSpace(req.Name),
		Aliases:                  pq.StringArray(aliases),
		Status:                   "ACTIVE",
		IngestionIntervalMinutes: req.IngestionIntervalMinutes,
	}

	if err := s.repo.CreateEntity(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) GetEntity(id string) (*Entity, error) {

	entityID, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return s.repo.GetEntity(entityID)
}

func (s *service) ListEntities() ([]Entity, error) {
	return s.repo.ListEntities()
}

func normalizeAliases(
	aliases []string,
) []string {

	result := make([]string, 0)

	seen := make(map[string]struct{})

	for _, alias := range aliases {

		normalized :=
			strings.TrimSpace(
				strings.ToLower(alias),
			)

		if normalized == "" {
			continue
		}

		if _, exists := seen[normalized]; exists {
			continue
		}

		seen[normalized] = struct{}{}

		result = append(result, normalized)
	}

	return result
}
