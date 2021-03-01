package app

import (
	"encoding/json"
	"errors"
	"github.com/netwar1994/httpserver/cmd/bank/app/dto"
	"github.com/netwar1994/httpserver/pkg/card"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cardSvc *card.Service
	mux *http.ServeMux
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux) *Server {
	return &Server{cardSvc: cardSvc, mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/addCard", s.addCard)
	//s.mux.HandleFunc("/editCard", s.editCard)
	//s.mux.HandleFunc("/removeCard", s.removeCard)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	extractId := r.URL.Query().Get("id")
	if extractId == "" {
		response(w, errors.New("user id unspecified"))
		return
	}
	userId, err := strconv.ParseInt(extractId, 10, 64)
	if err != nil {
		response(w, errors.New("user id unspecified "))
	}

	cards := s.cardSvc.GetCardsByUserId(userId)

	dtos := make([]*dto.CardDTO, len(cards))
	for i, c := range cards {
		if c.Id == userId {
			dtos[i] = &dto.CardDTO{
				Id:       c.Id,
				Issuer:   c.Issuer,
				Number:   c.Number,
				Currency: c.Currency,
				Virtual:  c.Virtual,
			}
		}
	}

	response(w, dtos)
}

func (s *Server) addCard(w http.ResponseWriter, r *http.Request) {
	var c dto.CardDTO
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.cardSvc.StoreCard(c.Id, c.Issuer, c.Currency, c.Virtual)
	if err != nil {
		log.Println(err)
		response(w, err.Error())
		return
	}
	log.Println("card added: ", c.Id, c.Issuer, c.Number, c.Currency, c.Virtual)
	response(w, c)
}

func response(w http.ResponseWriter, dtos... interface{}) {
	rBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(rBody)
	if err != nil {
		log.Println(err)
		return
	}
}
