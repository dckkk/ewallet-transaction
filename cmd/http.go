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
