package loan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncephamz/tunaiku-loan-service/src/domain/loan"
	model "github.com/ncephamz/tunaiku-loan-service/src/domain/loan"
	"github.com/ncephamz/tunaiku-loan-service/src/infrastructure/responses"
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
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = loan.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.repository.CreateLoan(&loan)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
