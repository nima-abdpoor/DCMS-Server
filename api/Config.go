package api

import (
	db "DCMS/db/sqlc"
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigResponse struct {
	RequestUrls []string      `json:"validRequestUrls"`
	UrlIdFirst  []int64       `json:"urlIdFirst"`
	UrlIdSecond []UrlIdSecond `json:"urlIdSecond"`
	IsLive      bool          `json:"isLive"`
	SyncType    string        `json:"syncType"`
}

type UrlIdSecond struct {
	UrlId       int64  `json:"urlId"`
	Regex       string `json:"regex"`
	StartIndex  int32  `json:"startIndex"`
	FinishIndex int32  `json:"finishIndex"`
}
type getConfigRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getConfig(ctx *gin.Context) {
	var req getConfigRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.ID <= 10 {
		ctx.JSON(http.StatusForbidden, "This ID Is Reserved")
		return
	}
	result, err := server.store.GetConfigTx(context.Background(), db.GetConfigTxParams{ID: req.ID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) getDefaultConfig(ctx *gin.Context) {
	configResult, err := server.store.GetConfigTx(context.Background(), db.GetConfigTxParams{ID: util.DefaultConfigId})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, configResult)
}
