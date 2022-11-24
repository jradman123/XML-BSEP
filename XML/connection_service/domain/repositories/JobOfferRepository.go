package repositories

import (
	"connection/module/domain/model"
	"context"
)

type JobOfferRepository interface {
	Create(m *model.JobOffer, ctx context.Context) (model.JobOffer, error)
}
