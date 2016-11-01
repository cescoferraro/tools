package web

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
	"github.com/cescoferraro/tools/logger"
)


var toolsLogger = logger.New("TOOLS")

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


//EnableCORS is a middleware
func WideOpenCORS() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers",
					"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			}
			// Stop here if its Preflighted OPTIONS request
			if r.Method == "OPTIONS" {
				toolsLogger.Print("CORS preflight request!")
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}



type Module interface {
	SetRoutes(*mux.Router) *mux.Router
	Run()
}

type Modules []Module

func (modules Modules) Init() *mux.Router{
	baseRouter := mux.NewRouter()
	for _, a := range modules {
		baseRouter = a.SetRoutes(baseRouter)
		a.Run()
	}
	return baseRouter
}
