package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"simplebank/common"
	db "simplebank/db/sqlc"
	"strconv"
)

// GetAccount get one account with specified id.
func (server *Server) GetAccount(c *gin.Context) {
	para := c.Param("id")
	id, err := strconv.ParseInt(para, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidRequestParam))
		return
	}
	account, err := server.store.GetAccount(c, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorWithCode(common.AccountNotFound))
			return
		}
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, Success(account))
}

type listAllAccountsRequest struct {
	// Use `binding` tag to validate.
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListAllAccounts list all accounts organized by page.
func (server *Server) ListAllAccounts(c *gin.Context) {
	var req listAllAccountsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidRequestParam))
		return
	}
	arg := db.ListAllAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAllAccounts(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, ErrorWithCode(common.AccountNotFound))
		return
	}
	c.JSON(http.StatusOK, Success(accounts))
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR JPY CNY GBP"`
}

// CreateAccount create a new account with initial balance 0.
func (server *Server) CreateAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidRequestBody))
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(c, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Error(err))
			return
		}
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, Success(account))
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListAccounts list all accounts owned by one user.
func (server *Server) ListAccounts(c *gin.Context) {
	var req listAccountsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidRequestParam))
		return
	}
	owner := c.Param("owner")
	arg := db.ListAccountsParams{
		Owner:  owner,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, ErrorWithCode(common.AccountNotFound))
		return
	}
	c.JSON(http.StatusOK, Success(accounts))
}
