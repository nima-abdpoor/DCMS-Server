package api

import (
	"DCMS/db/postgresql/sqlc"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type LiveLog struct {
	Request struct {
		URL         string `json:"url"`
		Body        string `json:"body"`
		Header      string `json:"header"`
		RequestTime string `json:"requestTime"`
	} `json:"request"`
	Response struct {
		Body        string `json:"body"`
		Code        string `json:"code"`
		Header      string `json:"header"`
		RequestTime string `json:"requestTime"`
	} `json:"response"`
}

func (server *Server) postLog(ctx *gin.Context) {
	var idPath idPath
	if err := ctx.ShouldBindUri(&idPath); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if idPath.ID <= 10 {
		ctx.JSON(http.StatusForbidden, "This ID Is Reserved")
		return
	}
	_, err := server.store.GetConfigTx(context.Background(), db.GetConfigTxParams{ID: idPath.ID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if err = os.MkdirAll("LiveLogs/"+strconv.FormatInt(idPath.ID, 10), os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	f, err := os.Create("LiveLogs/" + strconv.FormatInt(idPath.ID, 10) + "/" + strconv.FormatInt(idPath.ID, 10) + ".txt")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer f.Close()
	length := ctx.Request.Header.Get("Content-Length")
	intle, _ := strconv.Atoi(length)
	bodyBuffer := make([]byte, intle*2)
	body, err := ctx.Request.Body.Read(bodyBuffer)
	rawBody := string(bodyBuffer[0:body])
	fmt.Println("rawBody:", rawBody)
	_, err = f.WriteString(rawBody)
	fmt.Println("length:", intle)
	askdf, _ := ctx.GetRawData()
	fmt.Println("getrawData:", string(askdf))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Ok!")
}
