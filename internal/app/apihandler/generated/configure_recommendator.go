// This file is safe to edit. Once it exists it will not be overwritten

package generated

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
)

//go:generate swagger generate server --target ../../apihandler --name Recommendator --spec ../../../../api/swagger.yaml --api-package specops --model-package generated/specmodels --server-package generated --principal interface{} --exclude-main

func configureFlags(api *specops.RecommendatorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *specops.RecommendatorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// specops.PostRecommendationsMaxParseMemory = 32 << 20
	service := apihandler.NewRecommendatorService()
	api.GetRecommendationsBarcodeHandler = specops.GetRecommendationsBarcodeHandlerFunc(service.GetRecommendationsBarcodeHandler)
	api.PostRecommendationsHandler = specops.PostRecommendationsHandlerFunc(service.PostRecommendationsHandler)
	api.GetPingHandler = specops.GetPingHandlerFunc(service.GetPingHandler)

	if api.GetRecommendationsBarcodeHandler == nil {
		api.GetRecommendationsBarcodeHandler = specops.GetRecommendationsBarcodeHandlerFunc(func(params specops.GetRecommendationsBarcodeParams) middleware.Responder {
			return middleware.NotImplemented("operation specops.GetRecommendationsBarcode has not yet been implemented")
		})
	}
	if api.PostRecommendationsHandler == nil {
		api.PostRecommendationsHandler = specops.PostRecommendationsHandlerFunc(func(params specops.PostRecommendationsParams) middleware.Responder {
			return middleware.NotImplemented("operation specops.PostRecommendations has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	s.ReadTimeout = time.Minute * 2
	s.WriteTimeout = time.Minute * 2
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
