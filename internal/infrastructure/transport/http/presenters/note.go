package presenters

import (
	"github.com/gin-gonic/gin"
	openapi "golang-clean-architecture/api/http"
	"golang-clean-architecture/internal/app/domain"
	"golang-clean-architecture/internal/app/ports/out/dao/dto"
	"golang-clean-architecture/pkg/utils"
	"time"
)

func PresentToResp(note *domain.Note, context *gin.Context) any {
	nowString := time.Now().String()
	return openapi.SingleNoteBackendResponse{
		Payload: &openapi.NoteResponse{
			Description: &note.Description,
			Id:          &note.ID,
			Link:        &note.Link,
			Name:        &note.Name,
		},
		Meta: &openapi.MetaData{
			Path:      &context.Request.URL.Path,
			Timestamp: &nowString,
		},
	}
}

func PresentToReq(req openapi.CreateNoteRequest) dto.CreateNoteDto {
	return dto.CreateNoteDto{
		Description: utils.GetStringOrDefault(req.Description, ""),
		Link:        utils.GetStringOrDefault(req.Link, ""),
		Name:        utils.GetStringOrDefault(req.Name, ""),
	}
}
