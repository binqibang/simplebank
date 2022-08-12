package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"simplebank/common"
	db "simplebank/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(c *gin.Context) {
	var req transferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidRequestBody))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	_, valid := server.validAccount(c, req.FromAccountID, req.Currency, req.Amount)
	if !valid {
		return
	}

	result, err := server.store.TransferTx(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}

	c.JSON(http.StatusOK, Success(result))

}

func (server *Server) validAccount(c *gin.Context, accountID int64, currency string, amount int64) (db.Account, bool) {
	account, err := server.store.GetAccount(c, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorWithCode(common.AccountNotFound))
			return account, false
		}

		c.JSON(http.StatusInternalServerError, Error(err))
		return account, false
	}

	// Different currency.
	if account.Currency != currency {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidCurrency))
		return account, false
	}

	// Insufficient balance.
	if account.Balance < amount {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InsufficientBalance))
		return account, false
	}

	return account, true
}
