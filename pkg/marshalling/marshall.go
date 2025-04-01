package marshalling

import (
	"github.com/gin-gonic/gin"
	"golang-clean-architecture/pkg/errors"
	"net/http"
)

func HandleMarshalling[T any](c *gin.Context, req *T) error {
	if err := c.ShouldBindJSON(req); err != nil {
		errors.SetResponseError(c, errors.MarshallError)
		return err
	}
	return nil
}

func HandleResponse[T any](c *gin.Context, logic func() (*T, error), present func(*T, *gin.Context) any) {
	if result, err := logic(); err != nil {
		errors.HandleError(c, err)
	} else {
		c.JSON(http.StatusOK, present(result, c))
	}
}
