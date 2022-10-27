package app

import (
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
	"github.com/go-openapi/runtime/middleware"
)

type Service struct{}

func (srv *Service) PostRecommendationsHandler(params specops.PostRecommendationsParams, principal *interface{}) middleware.Responder {
	return middleware.NotImplemented("operation specops PostRecommendations has not yet been implemented")
}
