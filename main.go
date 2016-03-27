package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/saromanov/company-service/db"
	"github.com/saromanov/company-service/handler"
	proto "github.com/saromanov/company-service/proto/company"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.company"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/user",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()
	db.Init()

	proto.RegisterAccountHandler(service.Server(), new(handler.Account))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
