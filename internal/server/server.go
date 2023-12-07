package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Gggdevv/go-web-template/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server

	engine *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	engine := gin.Default()
	addr := fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port)
	return &Server{
		engine: engine,
		srv: &http.Server{
			Addr:    addr,
			Handler: engine,
		},
	}
}

func (server *Server) Start() error {
	group := server.engine.Group("/gorum")
	server.setupApis(group)

	if err := server.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
		return err
	}

	return nil
}

func (server *Server) ShutDown(ctx context.Context) error {
	return server.srv.Shutdown(ctx)
}
