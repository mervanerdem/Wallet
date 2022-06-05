package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"wallet/Model"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func NewServer(storage Model.WStorage) http.Handler {
	router := gin.New()
	router.Use(getInfoLogger(getLogger("./log.log")))
	//Get balance
	router.GET("/api/v1/wallets/:id/balance", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "Unsuitable ID number",
			})
			return
		}

		wallet, err := storage.Get(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]any{
			"Balance":   wallet.Wallet_balance,
			"Wallet_ID": wallet.ID,
		})

	})

	//Send debit
	router.POST("/api/v1/wallets/:id/debit", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "Unsuitable ID number",
			})
			return
		}
		var data = struct {
			Amount decimal.Decimal
		}{}
		ctx.BindJSON(&data)

		wallet, err := storage.Get(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		err = wallet.Debit(data.Amount)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{
				"Debit": "Amount can not be negative",
			})
			return
		}
		ctx.JSON(201, map[string]any{
			"Balance":   wallet.Wallet_balance,
			"Debit":     data.Amount,
			"Wallet_ID": wallet.ID,
		})
		storage.Update(Model.Wallet{
			ID:             wallet.ID,
			Wallet_balance: wallet.Wallet_balance,
		})
	})
	//Send Credit
	router.POST("/api/v1/wallets/:id/credit", func(ctx *gin.Context) {
		id_str := ctx.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "Unsuitable ID number",
			})
			return
		}
		var data = struct {
			Amount decimal.Decimal //body
		}{}
		ctx.BindJSON(&data)

		wallet, err := storage.Get(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = wallet.Credit(data.Amount)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{
				"Credit": "Amount Unsuitable",
			})
			return
		}
		ctx.JSON(201, map[string]any{
			"Balance":   wallet.Wallet_balance,
			"Debit":     data.Amount,
			"Wallet_ID": wallet.ID,
		})
		storage.Update(Model.Wallet{
			ID:             wallet.ID,
			Wallet_balance: wallet.Wallet_balance,
		})

	})

	return router
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
