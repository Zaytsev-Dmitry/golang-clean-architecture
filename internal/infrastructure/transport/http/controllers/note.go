package controllers

import (
	"github.com/gin-gonic/gin"
	openapi "golang-clean-architecture/api/http"
	"golang-clean-architecture/internal/app/domain"
	"golang-clean-architecture/internal/app/ports/in/delegate"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	"golang-clean-architecture/internal/infrastructure/transport/http/presenters"
	"golang-clean-architecture/pkg"
)

type NoteController struct {
	delegate *delegate.NoteDelegate
}

func (cntr *NoteController) SaveNote(c *gin.Context) {
	var req openapi.CreateNoteRequest

	if err := pkg.HandleMarshalling(c, &req); err != nil {
		return
	}

	pkg.HandleResponse(c, func() (*domain.Note, error) {
		return cntr.delegate.CreateAndSave(presenters.PresentToReq(req))
	}, presenters.PresentToResp)
}

func (cntr *NoteController) GetNotesById(c *gin.Context, id int64) {
	pkg.HandleResponse(c, func() (*domain.Note, error) {
		return cntr.delegate.GetById(id)
	}, presenters.PresentToResp)
}

func Create(dao *noteDao.Dao) *NoteController {
	return &NoteController{
		delegate: delegate.Create(dao),
	}
}
