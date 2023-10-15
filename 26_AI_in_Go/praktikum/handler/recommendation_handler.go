package handler

import (
	"ai_in_go/helper"
	"ai_in_go/model/web"
	"ai_in_go/service"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RecommendationHandler interface {
	GetRecommendations(writer http.ResponseWriter, request *http.Request, params httprouter.Params) 
}

type RecommendationHandlerImpl struct {
	service service.RecommendationService
}

func NewRecommendationHandler(service service.RecommendationService) RecommendationHandler {
	return &RecommendationHandlerImpl{service: service}
}

func (handler *RecommendationHandlerImpl) GetRecommendations(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	recommendationRequest := web.RecommendationRequest{}

	err := helper.ReadFromRequestBody(request, &recommendationRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: err.Error(),
			Data: nil,
		})
	}

	recommendationMessage := fmt.Sprintf("I'm looking for a laptop recommendation. My budget is IDR%d, and I plan to use it for %s. I prefer a laptop with %s operating system. What would you recommend?", recommendationRequest.Budget, recommendationRequest.Purpose, recommendationRequest.OS)

	response, err := handler.service.GetRecommendations(recommendationMessage)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code: http.StatusInternalServerError,
			Status: err.Error(),
			Data: nil,
		})
	}

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Data: response,
	})
}