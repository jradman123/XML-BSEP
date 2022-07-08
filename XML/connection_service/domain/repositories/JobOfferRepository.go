package repositories

import "connection/module/domain/model"

type JobOfferRepository interface {
	Create(m *model.JobOffer) (model.JobOffer, error)
}
