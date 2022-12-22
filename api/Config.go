package api

import (
	"DCMS/db/sqlc"
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigResponse struct {
	RequestUrls  []string      `json:"validRequestUrls"`
	UrlIdFirst   []int64       `json:"urlIdFirst"`
	UrlIdSecond  []UrlIdSecond `json:"urlIdSecond"`
	IsLive       bool          `json:"isLive"`
	SyncType     string        `json:"syncType"`
	SaveRequest  bool          `json:"saveRequest"`
	SaveResponse bool          `json:"saveResponse"`
	SaveError    bool          `json:"saveError"`
	SaveSuccess  bool          `json:"saveSuccess"`
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

type PostConfig struct {
	ID           int64       `form:"id" json:"id" binding:"required"`
	SyncType     string      `form:"sync_type" json:"sync_type" binding:"required"`
	IsLive       bool        `form:"is_live" json:"is_live"`
	SaveRequest  bool        `form:"saveRequest" json:"saveRequest"`
	SaveResponse bool        `form:"saveResponse" json:"saveResponse"`
	SaveError    bool        `form:"saveError" json:"saveError"`
	SaveSuccess  bool        `form:"saveSuccess" json:"saveSuccess"`
	UrlHashFirst []string    `form:"urlHashFirst" json:"urlHashFirst" binding:"required"`
	UrlSecond    []UrlSecond `form:"urlSecond" json:"urlSecond" binding:"required"`
	RequestUrl   []string    `form:"requestUrl" json:"requestUrl" binding:"required"`
}

type UrlSecond struct {
	UrlHash string  `form:"url_hash" json:"url_hash" binding:"required"`
	Regex   []Regex `form:"regex" json:"regex" binding:"required"`
}

type Regex struct {
	Regex       string `form:"regex" json:"regex" binding:"required"`
	StartIndex  int32  `form:"start_index" json:"start_index" binding:"required"`
	FinishIndex int32  `form:"finish_index" json:"finish_index" binding:"required"`
}

func (server *Server) postConfig(ctx *gin.Context) {
	var req PostConfig
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	configResult, err := server.store.AddConfigTx(context.Background(), db.AddConfigTxParams{
		ID:           req.ID,
		SyncType:     req.SyncType,
		IsLive:       req.IsLive,
		SaveRequest:  req.SaveRequest,
		SaveResponse: req.SaveResponse,
		SaveError:    req.SaveError,
		SaveSuccess:  req.SaveSuccess,
		UrlHashFirst: req.UrlHashFirst,
		UrlSecond:    MapToUrlSecondDb(req.UrlSecond, req.ID),
		RequestUrl:   req.RequestUrl,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, configResult)
}
