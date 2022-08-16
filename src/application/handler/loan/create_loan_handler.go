package loan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncephamz/tunaiku-loan-service/src/domain/loan"
	model "github.com/ncephamz/tunaiku-loan-service/src/domain/loan"
)

type CreateLoanHandler struct {
	repository loan.LoanRepository
}

func NewCreateLoanHandler(repository loan.LoanRepository) *CreateLoanHandler {
	return &CreateLoanHandler{
		repository: repository,
	}
}

func (h *CreateLoanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loan := model.Loan{}

	err := json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = loan.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.repository.CreateLoan(&loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
