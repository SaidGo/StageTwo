//go:build disable_extras

package gates

import (
	"fmt"

	"example.com/local/Go2part/domain"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func (a *Service) ProjectCreate(project domain.Project, userUUID uuid.UUID) error {
	fUUIDs := a.dict.GetUserFederatons(userUUID.String())

	hasFederation := lo.IndexOf(fUUIDs, project.FederationUUID)

	if hasFederation == -1 {
		return fmt.Errorf("федерация не найдена")
	}

	companyDTO, found := a.dict.FindCompany(project.CompanyUUID)

	if !found {
		return fmt.Errorf("компании не существует")
	}

	if companyDTO.ProjectsTotal >= a.projectLimits {
		return fmt.Errorf("превышен лимит проектов (максимум %v)", a.projectLimits)
	}

	return nil
}

func (a *Service) ProjectDelete(project domain.Project, userUUID uuid.UUID) error {
	cUUIDs := a.dict.GetUserCompanies(userUUID.String())

	hasCompany := lo.IndexOf(cUUIDs, project.CompanyUUID)

	if hasCompany == -1 {
		return fmt.Errorf("компания не найдена")
	}

	return nil
}

func (a *Service) ProjectPatch(project domain.Project, userUUID uuid.UUID) error {
	cUUIDs := a.dict.GetUserCompanies(userUUID.String())

	hasCompany := lo.IndexOf(cUUIDs, project.CompanyUUID)

	if hasCompany == -1 {
		return fmt.Errorf("компания не найдена")
	}

	return nil
}
