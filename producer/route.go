package producer

import (
	"github.com/Shopify/sarama"
	"github.com/astomiwanda/go-coinbit/lib"
	"github.com/astomiwanda/go-coinbit/producer/repository"
	"github.com/astomiwanda/go-coinbit/producer/service"
	"github.com/gorilla/mux"
	"net/http"
)

func RestRouter(producers sarama.SyncProducer) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	customerRouter(api, producers)
	return r
}

func customerRouter(r *mux.Router, producers sarama.SyncProducer) {
	kafka := KafkaProducer{
		Producer: producers,
	}

	var db = lib.ConnectionDB()
	var balanceRepo = repository.NewBalanceRepository(db)
	var balanceService = service.NewBalanceService(balanceRepo)
	var balanceHandler = NewBalanceHandler(kafka, balanceService)

	r.HandleFunc("/balance/{id}", balanceHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/balance", balanceHandler.Post).Methods(http.MethodPost)
}
