package main

import (
	"ai_in_go/app"
	"ai_in_go/config"
	"ai_in_go/handler"
	"ai_in_go/service"
	"net/http"
)

func main() {
	openAPIKey, err := config.NewEnvConfig()
	if err != nil {
		panic(err.Error())
	}
	
	recommendationService := service.NewRecommendationService(openAPIKey)
	recommendationHandler := handler.NewRecommendationHandler(recommendationService)

	router := app.NewRouter(recommendationHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}

}