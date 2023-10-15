package app

import (
	"ai_in_go/handler"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(recommendationHandler handler.RecommendationHandler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/recommendations", recommendationHandler.GetRecommendations)

	return router
}