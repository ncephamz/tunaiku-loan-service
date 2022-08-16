package loan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncephamz/tunaiku-loan-service/src/domain/loan"

	"github.com/gorilla/mux"
)

type LoanByIdHandler struct {
	repository loan.LoanRepository
}

func NewLoanByIdHandler(repository loan.LoanRepository) *LoanByIdHandler {
	return &LoanByIdHandler{
		repository: repository,
	}
}

func (h *LoanByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	response, err := h.repository.FindLoanById(id)
	if err != nil {
		fmt.Fprintf(w, string(err.Error()))
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, string(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
