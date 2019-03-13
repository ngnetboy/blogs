package main

import (
	"controller"
	"model"
	"net/http"
	"os"
	"os/signal"
	"service"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func init() {
	model.LoadConf()
	gin.SetMode(gin.ReleaseMode)
}

func signalHandler(ln *http.Server) {
	chSigInt := make(chan os.Signal, 1)
	signal.Notify(chSigInt, os.Signal(syscall.SIGINT))
	chSigTerm := make(chan os.Signal, 1)
	signal.Notify(chSigTerm, os.Signal(syscall.SIGTERM))

	log.Infoln("signal handler start")

	for {
		select {
		case <-chSigInt:
			log.Infoln("signal SIGINT")
			ln.Close()
			service.DisconnectDB()

		case <-chSigTerm:
			log.Infoln("signal SIGTERM")
			ln.Close()
			service.DisconnectDB()
		}
	}
}

func main() {
	log.Infoln("api server listen address ", model.Conf.Addr)

	router := controller.GetRouter()

	server := &http.Server{
		Addr:    model.Conf.Addr,
		Handler: router,
	}
	service.ConnectDB()
	go signalHandler(server)
	server.ListenAndServe()
}
