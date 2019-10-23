package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"RPN/api"
	"RPN/stdlog"

	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
)

var options struct {
	GraceTimeOut       time.Duration `long:"graceful-timeout" description:"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m" default:"15s"`
	ServerWriteTimeout time.Duration `long:"swtimeout" description:"the duration for server write time out" default:"15s"`
	ServerReadTimeout  time.Duration `long:"srtimeout" description:"the duration for server read timeout" default:"15s"`
	IPAdrress          string        `long:"ip" description:"IP address of service" default:"0.0.0.0" env:"IP_ADDRESS"`
	Port               string        `long:"port" description:"Port number of service" default:"8080" env:"PORT"`
	RPNLog             string        `long:"stdlog" description:"std log file path" default:"rpn_std.log" env:"STD_LOG"`
	RPNErrLog          string        `long:"errlog" description:"std error log file path" default:"rpn_error.log" env:"ERR_LOG"`
}

func main() {
	flags.Parse(&options)

	stdlog.SetLogFiles(options.RPNLog, options.RPNErrLog)
	logger := log.New(stdlog.Out, "main:", log.LstdFlags)
	logger.Println("starting server now")

	router := RegisterService()

	srv := &http.Server{
		Addr:         options.IPAdrress + ":" + options.Port,
		WriteTimeout: options.ServerWriteTimeout,
		ReadTimeout:  options.ServerReadTimeout,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), options.GraceTimeOut)
	defer cancel()
	srv.Shutdown(ctx)

	logger.Println("shutting down")
	os.Exit(0)
}

// RegisterService creates mux router
func RegisterService() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/parse", api.RpnHandler).Methods("POST")
	router.HandleFunc("/health", api.HealthHandler).Methods("GET")
	return router
}
