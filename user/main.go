package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/opentreehole/go-common"
	"github.com/ppolariss/paper-submission-review-system-go/user/api"
	"github.com/ppolariss/paper-submission-review-system-go/user/config"
	"github.com/ppolariss/paper-submission-review-system-go/user/model"
	"github.com/ppolariss/paper-submission-review-system-go/user/rpc"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init()
	model.Init()
	rpc.Init()

	var disableStartupMessage = false
	if config.Config.Mode == "prod" {
		disableStartupMessage = true
	}
	app := fiber.New(fiber.Config{
		ErrorHandler:          common.ErrorHandler,
		DisableStartupMessage: disableStartupMessage,
	})
	registerMiddlewares(app)
	api.RegisterRoutes(app)
	go func() {
		err := app.Listen("0.0.0.0:8000")
		if err != nil {
			log.Fatal().Err(err).Msg("app listen failed")
		}
	}()

	interrupt := make(chan os.Signal, 1)

	// wait for CTRL-C interrupt
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	// close app
	err := app.Shutdown()
	if err != nil {
		log.Err(err).Msg("error shutdown app")
	}
}

func registerMiddlewares(app *fiber.App) {
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(common.MiddlewareGetUserID)
	if config.Config.Mode != "bench" {
		app.Use(common.MiddlewareCustomLogger)
	}
	app.Use(pprof.New())
}
