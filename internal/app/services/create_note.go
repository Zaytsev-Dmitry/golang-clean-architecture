package services

import (
	"golang-clean-architecture/internal/app/domain"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	"golang-clean-architecture/internal/app/ports/out/dao/dto"
	"golang-clean-architecture/pkg/errors"
)

type CreateNoteUCase struct {
	dao *noteDao.Dao
}

func (c CreateNoteUCase) CreateAndSave(dto dto.CreateNoteDto) (*domain.Note, error) {
	save, err := c.dao.NoteRepo.Save(dto)
	if err != nil {
		return nil, errors.UpdateErrorText(err.Action, err.Err)
	}
	return save, nil
}

func NewCreateNoteUCase(dao *noteDao.Dao) *CreateNoteUCase {
	return &CreateNoteUCase{
		dao: dao,
	}
}
