package tools

import (
	"net/http"
	"os/signal"
	"syscall"
	"fmt"
	"strconv"
	"github.com/braintree/manners"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	"os"
	"log"
)

type Adapter func(http.Handler) http.Handler

// Adapt h with all specified adapters.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}


func GracefulStartServer(port int,router *mux.Router) {
	errChan := make(chan error, 10)
	go func() {
		toolsLogger.Print("Running the server at 0.0.0.0:"+strconv.Itoa(port))
		addr := "0.0.0.0:" + strconv.Itoa(port)
		handl := handlers.CombinedLoggingHandler(os.Stdout, router)
		errChan <- manners.ListenAndServe(addr, handl)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			manners.Close()
			os.Exit(0)
		}
	}
}