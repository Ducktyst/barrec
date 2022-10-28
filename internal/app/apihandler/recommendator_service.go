package apihandler

import (
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specmodels"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
	"github.com/go-openapi/runtime/middleware"
)

type RecommendatorService struct {
}

func NewRecommendatorService() *RecommendatorService {
	return &RecommendatorService{}
}

func (srv *RecommendatorService) GetRecommendationsBarcodeHandler(params specops.GetRecommendationsBarcodeParams) middleware.Responder {
	params.Barcode
	payload := make([]*specmodels.Recommendation, 0, 0)
	return specops.NewGetRecommendationsBarcodeOK().WithPayload(payload)
}

func (srv *RecommendatorService) PostRecommendationsHandler(params specops.PostRecommendationsParams) middleware.Responder {
	payload := make([]*specmodels.Recommendation, 0, 0)
	return specops.NewPostRecommendationsOK().WithPayload(payload)
}
