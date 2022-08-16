package loan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncephamz/tunaiku-loan-service/src/domain/loan"

	"github.com/gorilla/mux"

	"github.com/ncephamz/tunaiku-loan-service/src/infrastructure/responses"
)

type ListLoanByKtpHandler struct {
	repository loan.LoanRepository
}

func NewListLoanByKtpHandler(repository loan.LoanRepository) *ListLoanByKtpHandler {
	return &ListLoanByKtpHandler{
		repository: repository,
	}
}

func (h *ListLoanByKtpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ktp := mux.Vars(r)["ktp"]

	response, err := h.repository.FindAllByKtp(ktp)
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
