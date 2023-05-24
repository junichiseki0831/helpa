package infra

import (
	domain "helpa/src/core/domain/userdm"
	"helpa/src/core/infra/datamodel"
	"helpa/src/support/smperr"
	"time"
)

func toDomainUser(dmUser *datamodel.User) (*domain.User, error) {
	const layout = "2006-01-02 15:04:05"
	createdAtTime, err := time.Parse(layout, dmUser.CreatedAt)
	if err != nil {
		return nil, smperr.BadRequest("failed to parse createdat")
	}
	updatedAtTime, err := time.Parse(layout, dmUser.UpdatedAt)
	if err != nil {
		return nil, smperr.BadRequest("failed to parse updatedat")
	}

	user, err := domain.Reconstruct(dmUser.ID, dmUser.Name, dmUser.Password, dmUser.Email, dmUser.Introduction, dmUser.Note, dmUser.Image, createdAtTime, updatedAtTime)
	if err != nil {
		return nil, err
	}

	return user, nil
}
