package infra

import (
	"errors"
	domain "helpa/src/core/domain/userdm"
	"helpa/src/core/infra/datamodel"
	"time"
)

func toDomainUser(dmUser *datamodel.User) (*domain.User, error) {
	const layout = "2006-01-02 15:04:05"
	createdAtTime, err := time.Parse(layout, dmUser.CreatedAt)
	if err != nil {
		return nil, errors.New("failed to parse createdat")
	}
	updatedAtTime, err := time.Parse(layout, dmUser.UpdatedAt)
	if err != nil {
		return nil, errors.New("failed to parse updatedat")
	}

	user, err := domain.Reconstruct(dmUser.ID, dmUser.Name, dmUser.Password, dmUser.Email, dmUser.Introduction, dmUser.Note, dmUser.Image, createdAtTime, updatedAtTime)
	if err != nil {
		return nil, errors.New("reconstruct processing failed")
	}

	return user, nil
}
