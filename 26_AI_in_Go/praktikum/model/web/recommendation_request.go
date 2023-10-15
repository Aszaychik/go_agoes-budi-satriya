package web

type RecommendationRequest struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
	OS      string `json:"os"`
}