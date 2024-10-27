package cmd

import (
	"ewallet-transaction/external"
	"ewallet-transaction/helpers"
	"ewallet-transaction/internal/api"
	"ewallet-transaction/internal/interfaces"
	"ewallet-transaction/internal/repository"
	"ewallet-transaction/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/health", d.HealthcheckAPI.HealthcheckHandlerHTTP)

	transactionV1 := r.Group("/transaction/v1")
	transactionV1.POST("/create", d.MiddlewareValidateToken, d.TransactionAPI.CreateTransaction)
	transactionV1.PUT("/update-status/:reference", d.MiddlewareValidateToken, d.TransactionAPI.UpdateStatusTransaction)
	transactionV1.GET("/", d.MiddlewareValidateToken, d.TransactionAPI.GetTransaction)
	transactionV1.GET("/:reference", d.MiddlewareValidateToken, d.TransactionAPI.GetTransactionDetail)
	transactionV1.POST("/refund", d.MiddlewareValidateToken, d.TransactionAPI.RefundTransaction)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckAPI
	External       interfaces.IExternal
	TransactionAPI interfaces.ITransactionAPI
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	external := &external.External{}

	trxRepo := &repository.TransactionRepo{
		DB: helpers.DB,
	}
	trxService := &services.TransactionService{
		TransactionRepo: trxRepo,
		External:        external,
	}
	trxAPI := &api.TransactionAPI{
		TransactionService: trxService,
	}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		External:       external,
		TransactionAPI: trxAPI,
	}
}
