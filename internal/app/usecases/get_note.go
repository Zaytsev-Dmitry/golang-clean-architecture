package usecases

import (
	"golang-clean-architecture/internal/app/domain"
)

type GetNote interface {
	GetById(id int64) (*domain.Note, error)
}
