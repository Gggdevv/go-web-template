package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Gggdevv/go-web-template/config"
	"github.com/Gggdevv/go-web-template/internal/datastore"
	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
	cfg *config.Config

	engine *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	engine := gin.Default()
	addr := fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port)
	return &Server{
		engine: engine,
		cfg:    cfg,
		srv: &http.Server{
			Addr:    addr,
			Handler: engine,
		},
	}
}

func (server *Server) Start() error {
	if err := datastore.Init(server.cfg); err != nil {
		return err
	}

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
