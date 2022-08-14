package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "simplebank/db/sqlc"
	"simplebank/util"
)

type Server struct {
	config util.Config
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	// Register custom validators.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil, err
		}
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	GroupV1 := router.Group("/v1")
	{
		GroupV1.GET("/account/:id", server.GetAccount)
		GroupV1.GET("/accounts/:owner", server.ListAccounts)
		GroupV1.GET("/accounts", server.ListAllAccounts)
		GroupV1.POST("/account", server.CreateAccount)

		GroupV1.POST("/transfer", server.createTransfer)

		GroupV1.POST("/user", server.CreateUser)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
