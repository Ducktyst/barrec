package apihandler

import (
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specmodels"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
	"github.com/ducktyst/bar_recomend/internal/barcode"
	"github.com/ducktyst/bar_recomend/internal/barcode/analyzer/common"
	"github.com/go-openapi/runtime/middleware"
)

type RecommendatorService struct {
}

func NewRecommendatorService() *RecommendatorService {
	return &RecommendatorService{}
}

func (srv *RecommendatorService) GetRecommendationsBarcodeHandler(params specops.GetRecommendationsBarcodeParams) middleware.Responder {
	articul, err := barcode.GetProductArticul(params.Barcode)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	kazanexpressRecommendation, err := common.GetPriceFrom(common.KazanExpress, articul)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	res := make([]common.Recommendation, 0, 5)
	res[0] = kazanexpressRecommendation

	payload := make([]*specmodels.Recommendation, 0, 1)
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul: res[i].Name,
			Barcode: params.Barcode,
			Price:   float64(res[i].Price),
			URL:     res[i].Url,
		}
	}

	return specops.NewGetRecommendationsBarcodeOK().WithPayload(payload)
}

func (srv *RecommendatorService) PostRecommendationsHandler(params specops.PostRecommendationsParams) middleware.Responder {

	if params.Content == nil { // possible?
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: "file is empty"})
	}

	img_barcode, err := barcode.ScanBarCodeFile(params.Content)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	articul, err := barcode.GetProductArticul(img_barcode)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	kazanexpressRecommendation, err := common.GetPriceFrom(common.KazanExpress, articul)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	res := make([]common.Recommendation, 0, 5)
	res[0] = kazanexpressRecommendation

	payload := make([]*specmodels.Recommendation, 0, 1)
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul: res[i].Name,
			Barcode: img_barcode,
			Price:   float64(res[i].Price),
			URL:     res[i].Url,
		}
	}
	return specops.NewPostRecommendationsOK().WithPayload(payload)
}
