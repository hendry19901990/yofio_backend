package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hendry19901990/yofio_backend/models"
	"github.com/hendry19901990/yofio_backend/repository"
	"github.com/hendry19901990/yofio_backend/services"
)

// success
// curl -X POST -d '{"investment": 1500}' http://localhost:9090/api/statistics

// unsuccess
//curl -X POST -d '{"investment": 1600}' http://localhost:9090/api/credit-assignment

func (cont *Controller) CreditAssignment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := models.InvestmentRequest{}
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, http.StatusText(400), 400)
		cont.WriteResponse(w, []byte(http.StatusText(400)), 400)
		return
	}

	conn, connErr := cont.GetConnection()
	if connErr != nil {
		cont.WriteResponse(w, []byte(connErr.Error()), 500)
		return
	}

	creditStore := repository.CreditStore{conn}

	c := services.GetCreditAssigner()
	creditType300, creditType500, creditType700, err := c.Assign(req.Investment)
	resp := models.CreditType{
		CreditType300: creditType300,
		CreditType500: creditType500,
		CreditType700: creditType700,
		Investment:    req.Investment,
	}

	if err != nil {
		creditStore.Save(&resp)
		fmt.Println(err)
		cont.WriteResponse(w, []byte(err.Error()), 400)
		return
	}

	resp.Success = 1
	creditStore.Save(&resp)

	bytesResp, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println(err)
		cont.WriteResponse(w, []byte(err.Error()), 500)
		return
	}

	cont.WriteResponse(w, bytesResp, 200)
}

//curl -X POST -d '{}' http://localhost:9090/api/statistics
func (cont *Controller) Statistics(w http.ResponseWriter, r *http.Request) {
	conn, connErr := cont.GetConnection()
	if connErr != nil {
		http.Error(w, connErr.Error(), 500)
		return
	}
	creditStore := repository.CreditStore{conn}
	statistics, err := creditStore.GetCreditTypeStatistics()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	bytesResp, err := json.Marshal(statistics)
	if err != nil {
		fmt.Println(err)
		cont.WriteResponse(w, []byte(err.Error()), 500)
		return
	}

	cont.WriteResponse(w, bytesResp, 200)
}
