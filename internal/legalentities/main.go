package legalentities

import "example.com/local/Go2part/domain"

type Service struct {
    repo Repository
}

func (s *Service) GetAll() ([]domain.LegalEntity, error) {
	ormEntities, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []domain.LegalEntity
	for _, orm := range ormEntities {
		result = append(result, orm.ToDomain())
	}
	return result, nil
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}