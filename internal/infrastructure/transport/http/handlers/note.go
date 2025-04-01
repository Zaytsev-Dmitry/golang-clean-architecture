package handlers

import (
	"github.com/gin-gonic/gin"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	"golang-clean-architecture/internal/infrastructure/transport/http/controllers"
)

type NoteBackendApi struct {
	controller *controllers.NoteController
}

func (cntr *NoteBackendApi) SaveNote(c *gin.Context) {
	cntr.controller.SaveNote(c)
}

func (cntr *NoteBackendApi) GetNotesById(c *gin.Context, id int64) {
	cntr.controller.GetNotesById(c, id)
}

func NewNoteBackendApi(dao *noteDao.Dao) *NoteBackendApi {
	return &NoteBackendApi{
		controller: controllers.Create(dao),
	}
}
