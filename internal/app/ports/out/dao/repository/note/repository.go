package note

import (
	"github.com/jmoiron/sqlx"
	"golang-clean-architecture/internal/app/domain"
	"golang-clean-architecture/internal/app/ports/out/dao/dto"
	"golang-clean-architecture/internal/app/ports/out/dao/queries"
	"golang-clean-architecture/pkg/errors"
	"strconv"
)

const (
	Get       = "Get"
	QueryRowx = "QueryRowx"
)

type Repository struct {
	db *sqlx.DB
}

func (n *Repository) Save(dto dto.CreateNoteDto) (*domain.Note, *errors.CustomError) {
	return n.executeQuery(
		QueryRowx,
		queries.InsertNoteQuery,
		"сохранение заметки",
		dto.Name, dto.Link, dto.Description,
	)
}

func (n *Repository) GetNoteById(id int64) (*domain.Note, *errors.CustomError) {
	return n.executeQuery(
		Get,
		queries.SelectNoteByID,
		"получении заметки по ID: "+strconv.FormatInt(id, 10),
		id,
	)
}

func (r *Repository) executeQuery(queryType string, query, action string, args ...interface{}) (*domain.Note, *errors.CustomError) {
	var err error
	var note domain.Note

	switch queryType {
	case Get:
		{
			err = r.db.Get(&note, query, args...)

		}
	case QueryRowx:
		{
			err = r.db.QueryRowx(query, args...).StructScan(&note)

		}
	}
	if err != nil {
		return nil, r.handleQueryError(err, action)
	}
	return &note, nil
}

func (n *Repository) handleQueryError(err error, action string) *errors.CustomError {
	return errors.New(action, err)
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
