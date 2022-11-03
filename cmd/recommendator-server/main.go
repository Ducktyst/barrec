package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
)

const port = 8091

func main() {

	swaggerSpec, err := loads.Embedded(generated.SwaggerJSON, generated.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := specops.NewRecommendatorAPI(swaggerSpec)
	server := generated.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "backend"
	parser.LongDescription = "# Introduction"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	server.Port = port
	server.ReadTimeout = time.Minute * 10
	server.WriteTimeout = time.Minute * 10

	// ngrok
	tun, err := ngrok.StartTunnel(context.Background(),
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("tunnel created:", tun.URL())

	// serve
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
