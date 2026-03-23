package handler

import (
	"encoding/json"
	"net/http"
	"url_shortener/logger"
	"url_shortener/model"
	"url_shortener/service"
	"github.com/gorilla/mux"
)

type ShortenerHandler struct {
	S *service.ShortenerService
	L *logger.Logger
}

func (s *ShortenerHandler) GenerateShortURL(rw http.ResponseWriter, req *http.Request) {
	s.L.LogMessage("Entering into handling routes for generating short url")
	var shortener_model model.Shortener_Model
	err := json.NewDecoder(req.Body).Decode(&shortener_model)
	if err != nil {
		s.L.LogFatalMessage("Failed to parse request body")
		http.Error(rw, "failed to parse request body", http.StatusInternalServerError)
	}
	s.S.CreateShortUrl(&shortener_model)
	json.NewEncoder(rw).Encode(&shortener_model)
}

func (s *ShortenerHandler) GetActualURLfromSHortURL(rw http.ResponseWriter, req *http.Request) {
	s.L.LogMessage("Entering into handling routes for returning actual url")
	id := mux.Vars(req)["id"]
	if id==""{
		s.L.LogFatalMessage("Failed to parse short url id")
		http.Error(rw, "failed to parse short url id", http.StatusInternalServerError)
		return
	}
	actual_url,err:=s.S.GetUrl(id)
	if err!=nil{
        s.L.LogFatalMessage("Failed to fetch from short url")
		http.Error(rw, "Failed to fetch from short url", http.StatusNotFound)
		return
	}
    http.Redirect(rw,req,string(actual_url),http.StatusMovedPermanently)
}
