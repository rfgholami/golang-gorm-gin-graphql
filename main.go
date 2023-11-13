package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/routes"
	"github.com/kwa0x2/GoLang-Postgre-API/service"
	"net/http"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	route := service.RoutePost()

	h := handler.New(&handler.Config{
		Schema:   &route,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":9898", nil)
	//router.Run(":9898")

}
