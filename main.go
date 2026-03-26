package main

import (
	"log"
	"os"
	"net/http"
	"url_shortener/handler"
	initializer "url_shortener/initializer"
	logger "url_shortener/logger"
    // "time"
	"github.com/gorilla/mux"
	"url_shortener/repository"
	"url_shortener/service"
	gohandlers "github.com/gorilla/handlers"
	"github.com/bits-and-blooms/bloom/v3"

)

//complex url --give it to simple value but should redirect it to original url

func main() {
	logr := logger.New(os.Stdout, "", log.LstdFlags)
	logr.LogMessage("Entering into url shortener main class")
	// initializer.Load_variables()
	dbi := initializer.Dbinitializer{L: logr}
	initializer.ConnectDb(&dbi)
	logr.LogMessage("DB is connected")
	rdbi := initializer.RedisClient{L: logr}
	rdb:=initializer.ConnectTORedis(&rdbi)
	logr.LogMessage("Redis is connected")
    repo1:= &repository.ShortenerRepository{DB: dbi.DB}
    srvc := &service.ShortenerService{Repo: repo1, L:logr , Rdb: rdb , Filter: bloom.NewWithEstimates(1000000, 0.001)}
    shrtn := &handler.ShortenerHandler{S: srvc, L: logr}
	srvmx := mux.NewRouter()
	Generateurl:=srvmx.Methods("POST").Subrouter()
	Generateurl.HandleFunc("/generate",shrtn.GenerateShortURL)
	Fetchurl:=srvmx.Methods("GET").Subrouter()
	Fetchurl.HandleFunc("/get_url/{id}",shrtn.GetActualURLfromSHortURL)
	ch:= gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	logr.LogMessage("going to http serverr")
	srv := &http.Server{
			Addr: ":8080",
			Handler: ch(srvmx),
			// IdleTimeout: 120*time.Second ,
			// ReadTimeout: 1*time.Second,
			// WriteTimeout: 1*time.Second,
		}
		logr.LogMessage("Server starting on port 8080")
		err := srv.ListenAndServe()
		if err != nil {
			logr.LogFatalMessage("Server failed to start")
		}
}
