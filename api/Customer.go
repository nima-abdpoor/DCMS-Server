package api

import (
	db "DCMS/db/postgresql/sqlc"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerPostResult struct {
	Username    string `json:"username"`
	Info        string `json:"info"`
	Email       string `json:"email"`
	PackageName string `json:"packageName"`
	SdkUuid     string `json:"sdk_uuid"`
}

func (server *Server) postCustomer(ctx *gin.Context) {
	var arg db.AddCustomerTxParams
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		fmt.Printf("here is the err %v\n", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	customerResult, err := server.Store.AddCustomerTx(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, CustomerPostResult{
		Username:    customerResult.Customer.Username,
		Info:        customerResult.Customer.Info,
		Email:       customerResult.Customer.Email,
		PackageName: customerResult.Customer.PackageName,
		SdkUuid:     customerResult.Customer.SdkUuid,
	})
}
