package rest

import (
	"backend/internal/config"
	"backend/internal/rest/gen"
	"backend/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config, services *service.Services) *http.Server {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))
	handler := newHandler(cfg, services)
	api := r.Group("/api/v1")

	strict := gen.NewStrictHandler(handler,
		[]gen.StrictMiddlewareFunc{
			//LogOperation,
		},
	)

	gen.RegisterHandlersWithOptions(api, strict,
		gen.GinServerOptions{
			//ErrorHandler: func(c *gin.Context, err error, status int) {
			//   c.JSON(status, gin.H{"error": err.Error()})
			//},
		},
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "8878"),
		Handler: r.Handler(),
	}
	log.Println(server.Addr)
	return &server
}
