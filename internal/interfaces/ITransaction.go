package interfaces

import (
	"context"
	"ewallet-transaction/internal/models"

	"github.com/gin-gonic/gin"
)

type ITransactionRepo interface {
	CreateTransaction(ctx context.Context, trx *models.Transaction) error
	UpdateStatusTransaction(context.Context, string, string, string) error
	GetTransactionByReference(context.Context, string, bool) (models.Transaction, error)
	GetTransaction(ctx context.Context, userID int) ([]models.Transaction, error)
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, req *models.Transaction) (models.CreateTransactionResponse, error)
	UpdateStatusTransaction(ctx context.Context, tokenData models.TokenData, req *models.UpdateStatusTransaction) error
	RefundTransaction(ctx context.Context, tokenData models.TokenData, req *models.RefundTransaction) (models.CreateTransactionResponse, error)
	GetTransactionDetail(ctx context.Context, reference string) (models.Transaction, error)
	GetTransaction(ctx context.Context, userID int) ([]models.Transaction, error)
}

type ITransactionAPI interface {
	CreateTransaction(c *gin.Context)
	UpdateStatusTransaction(c *gin.Context)
	RefundTransaction(c *gin.Context)
	GetTransaction(c *gin.Context)
	GetTransactionDetail(c *gin.Context)
}
