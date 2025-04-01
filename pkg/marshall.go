package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleMarshalling[T any](c *gin.Context, req *T) error {
	if err := c.ShouldBindJSON(req); err != nil {
		setResponseError(c, MarshallError)
		return err
	}
	return nil
}

func HandleResponse[T any](c *gin.Context, logic func() (*T, error), present func(*T, *gin.Context) any) {
	if result, err := logic(); err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, present(result, c))
	}
}
