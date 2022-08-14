package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"simplebank/common"
	db "simplebank/db/sqlc"
	"simplebank/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorWithCode(common.InvalidRequestBody))
		return
	}
	// Store the hashed password.
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err))
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		FullName: req.FullName,
		Email:    req.Email,
	}
	user, err := server.store.CreateUser(c, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				switch pqErr.Constraint {
				case "users_pkey":
					c.JSON(http.StatusForbidden, ErrorWithCode(common.InvalidUsername))
				case "users_email_key":
					c.JSON(http.StatusForbidden, ErrorWithCode(common.InvalidEmail))
				}
			}
			return
		}
		c.JSON(http.StatusInternalServerError, Error(err))
		return
	}
	c.JSON(http.StatusOK, Success(user))
}
