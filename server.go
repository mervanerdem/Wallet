package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewServer(storage WStorage) http.Handler {
	router := gin.New()

	//Get balance
	router.GET("/api/v1/wallets/:id/balance", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			notFound(err, ctx)
		}

		wallet, err := storage.Get(id)
		if err != nil {
			notFound(err, ctx)
		} else {
			ctx.JSON(http.StatusOK, map[string]int{
				"Balance":   wallet.Balance(),
				"Wallet_ID": wallet.id,
			})
		}
	})

	//Send debit
	router.POST("/api/v1/wallets/:id/debit", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			notFound(err, ctx)
		}
		var data = struct {
			Amount int
		}{}
		ctx.BindJSON(&data)

		wallet, err := storage.Get(id)
		if err != nil {
			notFound(err, ctx)
		}
		wallet.Debit(data.Amount)

		ctx.JSON(201, map[string]int{
			"Balance":   wallet.Balance(),
			"Debit":     data.Amount,
			"Wallet_ID": wallet.id,
		})
	})
	//Send Credit
	router.POST("/api/v1/wallets/:id/credit", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			notFound(err, ctx)
		}
		var data = struct {
			Amount int
		}{}
		ctx.BindJSON(&data)

		wallet, err := storage.Get(id)
		if err != nil {
			notFound(err, ctx)
		}
		wallet.Debit(data.Amount)

		ctx.JSON(201, map[string]int{
			"Balance":   wallet.Balance(),
			"Credit":    data.Amount,
			"Wallet_ID": wallet.id,
		})
	})

	return router
}

func notFound(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, map[string]string{
		"error": err.Error(),
	})
}
