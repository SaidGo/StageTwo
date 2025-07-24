package legalentities

import "example.com/local/Go2part/domain"

type MockRepo struct{}

func (m *MockRepo) GetAll() ([]domain.LegalEntity, error) {
    return []domain.LegalEntity{
        {
            UUID:    "123e4567-e89b-12d3-a456-426614174000",
            Name:    "Test Entity",
            INN:     "1234567890",
            KPP:     "987654321",
            OGRN:    "1027700132195",
            Address: "Test Address",
        },
    }, nil
}
