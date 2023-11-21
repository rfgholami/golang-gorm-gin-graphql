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
	config.Connect()
	route := service.RoutePost()

	h := handler.New(&handler.Config{
		Schema:   &route,
		Pretty:   true,
		GraphiQL: false,
	})
	service.Register()
	service.Find("go")

	router := gin.New()

	routes.HealthCheck(router)
	routes.UserManagement(router)
	routes.HomePageRoute(router)
	routes.Person(router)

	http.Handle("/", router)

	http.Handle("/graphql", h)
	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		return
	}

}
