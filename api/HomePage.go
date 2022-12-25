package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
)

func homePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func init() {
	if _, err := os.Stat("public/single"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/single", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
