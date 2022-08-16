package app

import (
	"fmt"
	"net/http"

	handler "github.com/ncephamz/tunaiku-loan-service/src/application/handler/loan"
	repository "github.com/ncephamz/tunaiku-loan-service/src/infrastructure/persistence/loan"

	steinHq "github.com/ncephamz/tunaiku-loan-service/src/infrastructure/stein-hq"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	stein := steinHq.NewSteinHqConnection()

	repository := repository.NewSteinHqLoanRepository(stein)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	loansRouter := r.PathPrefix("/loan").Subrouter()
	loansRouter.Handle("/create", handler.NewCreateLoanHandler(repository)).Methods("POST")
	loansRouter.Handle("/{ktp}", handler.NewListLoanByKtpHandler(repository))
	loansRouter.Handle("/detail/{id}", handler.NewLoanByIdHandler(repository))

	return r
}
