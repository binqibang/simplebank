package v1

import (
	"github.com/gin-gonic/gin"
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
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
