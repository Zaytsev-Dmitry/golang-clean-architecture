package services

import (
	"database/sql"
	"errors"
	"golang-clean-architecture/internal/app/domain"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	"golang-clean-architecture/pkg"
)

type GetNoteUCase struct {
	dao *noteDao.Dao
}

func (c GetNoteUCase) GetById(id int64) (*domain.Note, error) {
	save, err := c.dao.NoteRepo.GetNoteById(id)
	var wrappedErr error
	if err != nil {
		if errors.Is(err.Err, sql.ErrNoRows) {
			wrappedErr = pkg.RowNotFound
		}
		return nil, pkg.UpdateErrorText(err.Action, wrappedErr)
	}
	return save, nil
}

func NewGetNoteUCase(dao *noteDao.Dao) *GetNoteUCase {
	return &GetNoteUCase{
		dao: dao,
	}
}
