package api

import (
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

func (server *Server) getConfig(ctx *gin.Context) {
	var response = getConfigSampleResponse()
	ctx.JSON(http.StatusOK, response)
}

func getConfigSampleResponse() ConfigResponse {
	return ConfigResponse{
		RequestUrls: []string{"http://192.168.1.111:8080", "http://192.168.43.145:8080"},
		UrlIdFirst:  []int64{1254444, 258774111},
		UrlIdSecond: []UrlIdSecond{
			{
				UrlId:       125477,
				Regex:       "aksdlfja",
				StartIndex:  0,
				FinishIndex: 1,
			},
			{
				UrlId:       158777,
				Regex:       "askldjf",
				StartIndex:  10,
				FinishIndex: 22,
			},
		},
		IsLive:   false,
		SyncType: "1",
	}
}
