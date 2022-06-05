package producer

import (
	"encoding/json"
	"fmt"
	"github.com/astomiwanda/go-coinbit/consumer/model"
	"github.com/astomiwanda/go-coinbit/lib"
	"github.com/astomiwanda/go-coinbit/producer/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BalanceHandler struct {
	kafka   KafkaProducer
	service service.BalanceService
}

func NewBalanceHandler(kafka KafkaProducer, service service.BalanceService) *BalanceHandler {
	return &BalanceHandler{kafka: kafka, service: service}
}

func (s *BalanceHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	walletID := []byte(vars["id"])

	balance, err := s.service.GetBalance(string(walletID))
	if err != nil {
		logrus.Errorf("Error handler get balance : %v", err)

		errMsg := fmt.Sprintf("Data not found")
		w.WriteHeader(http.StatusNotFound)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	if balance == nil {
		logrus.Errorf("Error handler balance null")

		errMsg := fmt.Sprintf("Data not found")
		w.WriteHeader(http.StatusNotFound)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	lib.ResponseBuilder(w, balance)
}

func (s *BalanceHandler) Post(w http.ResponseWriter, r *http.Request) {

	var api = &model.BalanceAPI{}
	err := json.NewDecoder(r.Body).Decode(api)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	var balance model.Balance
	err = lib.Merge(api, &balance)
	if err != nil {
		errMsg := fmt.Sprintf("Request merger error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	balanceByte, err := json.Marshal(balance)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	err = s.kafka.SendMessage("deposits", lib.BalancePost.ToKey(), balanceByte)
	if err != nil {
		errMsg := fmt.Sprintf("Create customer error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		lib.ResponseBuilder(w, errMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	lib.ResponseBuilder(w, "customer created")
}
