package api

import (
	"DCMS/db/postgresql/sqlc"
	"DCMS/parser"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (server *Server) uploadSingleFile(ctx *gin.Context) {
	saveUploadedFile(ctx)
	configs, err := parsUploadedFile()
	if err != nil {
		log.Fatal("Error In Parsing Config File...", err)
		return
	}
	server.saveConfigToDataBase(configs, ctx)
}

func (server *Server) saveConfigToDataBase(configs []db.AddConfigTxParams, ctx *gin.Context) {
	for _, config := range configs {
		configResult, err := server.Store.AddConfigTx(context.Background(), config)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			continue
		}
		ctx.JSON(http.StatusOK, configResult)
	}
}

func parsUploadedFile() ([]db.AddConfigTxParams, error) {
	return parser.ReadFromFile()
}

func saveUploadedFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	filePath := "http://localhost:8080/images/single/" + originalFileName + fileExt
	out, err := os.Create("public/single/" + originalFileName + fileExt)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
	defer out.Close()
}
