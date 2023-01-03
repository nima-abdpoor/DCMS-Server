package api

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type LogFile struct {
	File *multipart.FileHeader `form:"log" binding:"required"`
}

func (server *Server) postLogFile(ctx *gin.Context) {
	var req LogFile
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := ctx.SaveUploadedFile(req.File, "assets/"+req.File.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
