package delegate

import (
	"golang-clean-architecture/internal/app/domain"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	"golang-clean-architecture/internal/app/ports/out/dao/dto"
	"golang-clean-architecture/internal/app/services"
)

type NoteDelegate struct {
	createNoteUCase *services.CreateNoteUCase
	getNoteUCase    *services.GetNoteUCase
}

func (d *NoteDelegate) CreateAndSave(dto dto.CreateNoteDto) (*domain.Note, error) {
	save, err := d.createNoteUCase.CreateAndSave(dto)
	return save, err
}

func (d *NoteDelegate) GetById(id int64) (*domain.Note, error) {
	save, err := d.getNoteUCase.GetById(id)
	return save, err
}

func Create(dao *noteDao.Dao) *NoteDelegate {
	return &NoteDelegate{
		createNoteUCase: services.NewCreateNoteUCase(dao),
		getNoteUCase:    services.NewGetNoteUCase(dao),
	}
}
