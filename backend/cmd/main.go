package main

import (
	"backend/internal/rest"
	"backend/internal/rest/gen"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), ErrorsToJSON())

	r.Use(ginmiddleware.OapiRequestValidatorWithOptions(spec, &ginmiddleware.Options{
		Options: openapi3filter.Options{AuthenticationFunc: authFunc},
	}))

	srv := rest.NewServer(workoutSvc, userSvc)

	strict := api.NewStrictHandler(srv, []api.StrictMiddlewareFunc{
		LogOperation,
	})

	gen.RegisterHandlersWithOptions(r, strict, api.GinServerOptions{
		// ошибки разбора path и query параметров
		ErrorHandler: func(c *gin.Context, err error, status int) {
			c.JSON(status, gin.H{"error": err.Error()})
		},
	})

	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) }) // ручные роуты рядом

	log.Fatal(r.Run(":8080"))
}
