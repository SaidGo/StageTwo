//go:build disable_extras

package emails

import (
	"example.com/local/Go2part/pkg/postgres"
)

type EmailRepository struct {
	gorm *postgres.GDB
}

func NewRepository(db *postgres.GDB) *EmailRepository {
	return &EmailRepository{
		gorm: db,
	}
}

func (r *EmailRepository) StoreEmail(email IMessage) (string, error) {
	mail := &Mail{
		Subject: email.GetSubject(),
		Text:    email.GetBody(),
	}

	err := r.gorm.DB.Create(&mail).Error
	if err != nil {
		return "", err
	}

	return mail.UUID, nil
}
