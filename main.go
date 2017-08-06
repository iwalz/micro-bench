package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/iwalz/bench/config"
	"github.com/iwalz/bench/handler"
	stress "github.com/iwalz/bench/proto/stress"
	_ "github.com/micro/go-plugins/registry/kubernetes"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.bench"),
		micro.Version("latest"),
		micro.Flags(
			cli.IntFlag{
				Name:   "count",
				Usage:  "Number of iteractions for stressig the used ressource",
				Value:  5000,
				EnvVar: "STRESS_COUNT",
			},
			cli.StringFlag{
				Name:   "user",
				Usage:  "Database username for mysql database",
				Value:  "",
				EnvVar: "MYSQL_USER",
			},
			cli.StringFlag{
				Name:   "password",
				Usage:  "Database password for mysql database",
				Value:  "",
				EnvVar: "MYSQL_PASSWORD",
			},
			cli.StringFlag{
				Name:   "database",
				Usage:  "Database endpoint for mysql database",
				Value:  "",
				EnvVar: "MYSQL_DB",
			},
			cli.StringFlag{
				Name:   "endpoint",
				Usage:  "Endpoint for mysql database",
				Value:  "",
				EnvVar: "MYSQL_ENDPOINT",
			}),
	)

	var conf config.Config

	// Initialise service
	service.Init(micro.Action(func(c *cli.Context) {
		conf = config.FromContext(c)
	}))

	// Register Handler
	h := handler.NewStressHandler(conf)
	stress.RegisterStressHandler(service.Server(), h)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
