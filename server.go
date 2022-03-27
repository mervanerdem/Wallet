package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewServer(storage WStorage) http.Handler {
	router := gin.New()

	router.GET("/api/v1/wallets/:id/balance", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, _ := strconv.Atoi(id_str)

		wallet, _ := storage.Get(id)
		ctx.JSON(http.StatusOK, map[string]int{
			"Balance": wallet.Balance(),
		})
	})

	return router
}
