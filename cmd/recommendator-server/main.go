package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
)

const port = 8091

var (
	Db     *sqlx.DB
	DBhost = "localhost"
	DBport = 5432
	DBuser = "aleksej"
	DBname = "recommendator"
)

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
	// init db start
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", DBhost, DBport, DBuser, DBname)
	Db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer func() {
		logrus.Info("db connection closing...")
		Db.Close()
	}()

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	// init db end

	server.ConfigureAPI(Db)
	// server.ConfigureDB(Db)
	server.Port = port
	server.ReadTimeout = time.Minute * 10
	server.WriteTimeout = time.Minute * 10

	// ngrok
	// tun, err := ngrok.StartTunnel(context.Background(),
	// 	config.HTTPEndpoint(),
	// 	ngrok.WithAuthtokenFromEnv(),
	// )
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	// log.Println("tunnel created:", tun.URL())

	// serve
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
	logrus.Info("server stop serving")
}
