package api

import (
	"DCMS/input"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (server *Server) homePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func uploadSingleFile(ctx *gin.Context) {
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
	configs, err := input.ReadFromFile(out)
	fmt.Println("here is configs:", configs)
	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
	defer out.Close()
}

func init() {
	if _, err := os.Stat("public/single"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/single", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	if _, err := os.Stat("public/multiple"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/multiple", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}

type Catlog struct {
	Product_id string `json: "product_id"`
	Quantity   int    `json: "quantity"`
}
