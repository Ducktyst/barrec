package apihandler

import (
	"fmt"
	"time"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specmodels"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
	"github.com/ducktyst/bar_recomend/internal/barcode/analyzer/common"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type RecommendatorService struct {
}

func NewRecommendatorService() *RecommendatorService {
	return &RecommendatorService{}
}

func (srv *RecommendatorService) GetRecommendationsBarcodeHandler(params specops.GetRecommendationsBarcodeParams) middleware.Responder {
	fmt.Println("GetRecommendationsBarcodeHandler", params.Barcode)

	articul, err := common.GetProductArticul(params.Barcode)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	res, err := findByArticul(articul)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul:  res[i].Name,
			ShopName: res[i].ShopName,
			Barcode:  params.Barcode,
			Price:    float64(res[i].Price),
			URL:      res[i].Url,
		}
	}

	return specops.NewGetRecommendationsBarcodeOK().WithPayload(payload)
}

func (srv *RecommendatorService) PostRecommendationsHandler(params specops.PostRecommendationsParams) middleware.Responder {
	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler start")

	if params.Content == nil { // possible?
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: "file is empty"})
	}

	img_barcode, err := common.ScanBarCodeFile(params.Content)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler ", img_barcode, err)
	articul, err := common.GetProductArticul(img_barcode)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	res, err := findByArticul(articul)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul:  res[i].Name,
			ShopName: res[i].ShopName,
			Barcode:  img_barcode,
			Price:    float64(res[i].Price),
			URL:      res[i].Url,
		}
	}
	return specops.NewPostRecommendationsOK().WithPayload(payload)
}

func (srv *RecommendatorService) GetPingHandler(params specops.GetPingParams) middleware.Responder {
	return specops.NewGetPingOK().WithPayload(&specmodels.Pong{Text: "service done!"})
}

func findByArticul(articul string) ([]common.Recommendation, error) {
	// kazanexpressRecommendation, err := common.GetPriceFrom(common.KazanExpress, articul)
	// if err != nil {
	// 	return nil, err
	// }
	ymRecommendation, err := common.GetPriceFrom(common.YandexMarket, articul)
	if err != nil {
		return nil, err
	}

	res := make([]common.Recommendation, 2)
	// res[0] = kazanexpressRecommendation
	res[1] = ymRecommendation
	return res, err
}
