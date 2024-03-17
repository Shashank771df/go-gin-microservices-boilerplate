package main

import (
	//Primero se importa config para la funcion init
	c "app/microservices/base_/config"
	//Luego log para que inicialice con init
	"app/microservices/base_/log"
	//Finalmente lib y routes para que inicialice con sus respectivos init
	"app/microservices/base_/lib"
	"app/microservices/base_/routes"

	//Finalmente cualquier paquete
	"app/core"
	"app/core/alert"
	"app/core/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	if c.Env.Micro.App.Stage == core.StageLive {
		gin.SetMode(gin.ReleaseMode)
	}

	apiName := c.Env.Micro.App.Name
	//Iniciando Echo
	e := gin.New()

	middleware := core.Middleware{}
	// recover
	e.Use(middleware.EchoRecovery())
	// add cors
	e.Use(middleware.AddCors())
	// Handler Error
	e.Use(middleware.ErrorsHandler())

	// registro de grupos y rutas por configuracion
	for _, group := range routes.AppRouting {
		eGroup := e.Group(group.Prefix, group.Middlewares...)

		for _, route := range group.Routes {
			if c.Env.Micro.App.Stage != core.StageLive {
				log.Log.Debug(logger.LogInfo{Key: apiName, Value: "Endpoint " + route.Method + " " + group.Prefix + route.Path})
			}
			eGroup.Handle(route.Method, route.Path, route.Handler)
		}
	}

	// iniciamos el servidor
	log.Log.Debug(logger.LogInfo{Key: apiName, Value: "Server START"})
	uri := fmt.Sprintf("%s:%d", c.Env.Micro.App.Address, c.Env.Micro.App.Port)
	log.Log.Debug(logger.LogInfo{Key: apiName, Value: "Server STARTED and RUNNING at: " + uri})
	go StartUp{}.Run()

	server := &http.Server{
		Addr:         uri,
		Handler:      e,
		ReadTimeout:  time.Duration(c.Env.Micro.App.ServerTimeout) * time.Minute,
		WriteTimeout: time.Duration(c.Env.Micro.App.ServerTimeout) * time.Minute,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			server.ErrorLog.Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	message := "Server is STARTING SHUTDOWN PROCESS"
	log.Log.Info(logger.LogInfo{Key: apiName, Value: message})

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Env.Micro.App.ExitTimeout)*time.Second)
	defer cancel()

	message = "Server is SHUTTING DOWN NOW"
	log.Log.Info(logger.LogInfo{Key: apiName, Value: message})

	if err := server.Shutdown(ctx); err != nil {
		message = fmt.Sprintf("Server was SHUTTED DOWN SUDDENLY - Error: %v", err.Error())
		lib.Alert.SendMessage(alert.AlertInfo{
			Key:   apiName,
			Value: message,
		})
		log.Log.Error(logger.LogInfo{Key: apiName, Value: message})
		return
	}

	message = "Server has been SHUTTED DOWN NICELLY"
	log.Log.Info(logger.LogInfo{Key: apiName, Value: message})
}
