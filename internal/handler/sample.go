package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
)

type (
	SampleArg struct {
		Name string `json:"name"`
	}
	SampleReply struct {
		Message string `json:"message"`
	}
)

func Sample(u *usecase.SampleUsecase) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input *SampleArg
		if err := c.ShouldBindJSON(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		param := usecase.NewSampleInput(input.Name)
		result, err := u.Run(c, param)
		if err != nil { // TODO: 共通化する
			if errors.IsPrecondition(err) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		reply := new(SampleReply)
		reply.FromOutput(result)
		c.JSON(http.StatusOK, reply)
	}
}

func (r *SampleReply) FromOutput(output *usecase.SampleOutput) {
	r.Message = output.Message
}
