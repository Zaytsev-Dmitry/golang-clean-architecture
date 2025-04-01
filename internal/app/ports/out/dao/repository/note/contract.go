package note

import (
	"golang-clean-architecture/internal/app/domain"
	"golang-clean-architecture/internal/app/ports/out/dao/dto"
)

type NoteContract interface {
	Save(dto dto.CreateNoteDto) (*domain.Note, error)
	GetNoteById(id int64) (*domain.Note, error)
}
