package api

import (
	db "DCMS/db/sqlc"
	"context"
	"database/sql"
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
	var req LiveLog
	var idPath idPath
	if err := ctx.ShouldBindUri(&idPath); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if idPath.ID <= 10 {
		ctx.JSON(http.StatusForbidden, "This ID Is Reserved")
		return
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
	_, err = f.WriteString("RequestTime:" + req.Response.RequestTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Ok!")
}
