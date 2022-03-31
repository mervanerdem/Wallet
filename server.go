package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func NewServer(storage WStorage) http.Handler {
	router := gin.New()
	router.Use(getInfoLogger(getLogger("./log.log")))
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
			ctx.JSON(http.StatusOK, map[string]any{
				"Balance":   wallet.Wallet_balance,
				"Wallet_ID": wallet.ID,
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
			Amount decimal.Decimal
		}{}
		ctx.BindJSON(&data)

		wallet, err := storage.Get(id)
		if err != nil {
			notFound(err, ctx)
		}
		wallet.Debit(data.Amount)

		ctx.JSON(201, map[string]any{
			"Balance":   wallet.Wallet_balance,
			"Debit":     data.Amount,
			"Wallet_ID": wallet.ID,
		})
	})
	//Send Credit
	router.POST("/api/v1/wallets/:id/credit", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			notFound(err, ctx)
			return
		}
		var data = struct {
			Amount decimal.Decimal //body
		}{}
		ctx.BindJSON(&data)

		wallet, err := storage.Get(id)
		if err != nil {
			notFound(err, ctx)
			return
		}
		wallet.Debit(data.Amount)

		ctx.JSON(201, map[string]any{
			"Balance":   wallet.Wallet_balance,
			"Credit":    data.Amount,
			"Wallet_ID": wallet.ID,
		})
	})

	return router
}

func notFound(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, map[string]string{
		"error": err.Error(),
	})
}

func getLogger(file string) *logrus.Logger {
	logger := logrus.New()

	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	logger.SetOutput(f)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC1123Z,
	})

	return logger
}

func getInfoLogger(logger *logrus.Logger) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		ctx.Next()

		// Stop timer
		duration := time.Since(start)

		logger.WithFields(map[string]any{
			"client_ip":  ctx.ClientIP(),
			"duration":   duration,
			"method":     ctx.Request.Method,
			"path":       ctx.Request.RequestURI,
			"status":     ctx.Writer.Status(),
			"referrer":   ctx.Request.Referer(),
			"request_id": ctx.Writer.Header().Get("Request-Id"),
		}).Info()
	}
}
