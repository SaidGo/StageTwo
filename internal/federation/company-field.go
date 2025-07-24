package federation

import (
	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/dto"
	"example.com/local/Go2part/internal/helpers"
	"github.com/google/uuid"
)

func (s *Service) CreateCompanyField(cf *domain.CompanyField) (items dto.CompanyFieldDTO, err error) {
	orm, err := s.repo.CreateCompanyField(cf)
	if err != nil {
		return items, err
	}

	return dto.CompanyFieldDTO{
		UUID:        orm.UUID,
		Name:        orm.Name,
		Description: orm.Description,
		DataType:    orm.DataType,
		Hash:        orm.Hash,
		Icon:        orm.Icon,
	}, err
}

func (s *Service) PutCompanyField(pf *domain.CompanyField) error {
	return s.repo.PutCompanyField(pf)
}

func (s *Service) GetProjectFields(uid uuid.UUID) (items []domain.CompanyField, err error) {
	orm, err := s.repo.GetProjectFields(uid)
	if err != nil {
		return items, err
	}

	// @todo
	items = helpers.Map[CompanyFields, domain.CompanyField](orm, func(item CompanyFields, index int) domain.CompanyField {
		return domain.CompanyField{
			UUID:               item.UUID,
			Hash:               item.Hash,
			Name:               item.Name,
			Description:        item.Description,
			Icon:               item.Icon,
			DataType:           domain.FieldDataType(item.DataType),
			CompanyUUID:        item.CompanyUUID,
			RequiredOnStatuses: item.RequiredOnStatuses,
			Style:              item.Style,
		}
	})

	return items, err
}

func (s *Service) DeleteCompanyField(uid uuid.UUID) (err error) {
	return s.repo.DeleteCompanyField(uid)
}

func (s *Service) GetCompanyFields(uid uuid.UUID) (dmns []domain.CompanyField, err error) {
	return s.repo.GetCompanyFields(uid)
}
