package services

import (
	"database/sql"
	"errors"
	"golang-clean-architecture/internal/app/domain"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	errors2 "golang-clean-architecture/pkg/errors"
)

type GetNoteUCase struct {
	dao *noteDao.Dao
}

func (c GetNoteUCase) GetById(id int64) (*domain.Note, error) {
	save, err := c.dao.NoteRepo.GetNoteById(id)
	var wrappedErr error
	if err != nil {
		if errors.Is(err.Err, sql.ErrNoRows) {
			wrappedErr = errors2.RowNotFound
		}
		return nil, errors2.UpdateErrorText(err.Action, wrappedErr)
	}
	return save, nil
}

func NewGetNoteUCase(dao *noteDao.Dao) *GetNoteUCase {
	return &GetNoteUCase{
		dao: dao,
	}
}
