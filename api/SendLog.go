package api

import (
	"DCMS/db/postgresql/sqlc"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	_, err := server.Store.GetConfigTx(context.Background(), db.GetConfigTxParams{ID: idPath.ID})
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
	f, err := os.OpenFile("LiveLogs/"+strconv.FormatInt(idPath.ID, 10)+"/"+strconv.FormatInt(idPath.ID, 10)+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer f.Close()
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	rawBody := string(jsonData)
	bodyStartIndex := strings.Index(rawBody, "{")
	bodyFinishIndex := strings.Index(rawBody[bodyStartIndex:], "--*****")
	_, err = f.WriteString(rawBody[bodyStartIndex : bodyFinishIndex+bodyStartIndex])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Ok!")
}
