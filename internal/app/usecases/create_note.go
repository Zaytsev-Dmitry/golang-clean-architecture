package usecases

import (
	"golang-clean-architecture/internal/app/domain"
	"golang-clean-architecture/internal/app/ports/out/dao/dto"
)

type CreateNote interface {
	CreateAndSave(dto dto.CreateNoteDto) (*domain.Note, error)
}
