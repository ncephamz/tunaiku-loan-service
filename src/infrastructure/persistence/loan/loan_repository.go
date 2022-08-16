package loan

import (
	stein "github.com/nasrul21/go-stein"
	"github.com/ncephamz/tunaiku-loan-service/src/domain/loan"
	model "github.com/ncephamz/tunaiku-loan-service/src/domain/loan"
)

type SteinHqLoanRepository struct {
	client *stein.Stein
}

func NewSteinHqLoanRepository(client *stein.Stein) *SteinHqLoanRepository {
	return &SteinHqLoanRepository{
		client: client,
	}
}

func (r *SteinHqLoanRepository) CreateLoan(loan *loan.Loan) (*loan.Loan, error) {
	loans := model.NewLoan(loan)

	_, _, err := r.client.Insert("Sheet1", loans)
	if err != nil {
		return nil, err
	}

	loan.ID = loans[0].ID
	return loan, nil
}

func (r *SteinHqLoanRepository) FindAllByKtp(ktp string) ([]*loan.Loan, error) {
	var (
		result   []*loan.Loan
		filtered []*loan.Loan
	)

	readOption := stein.ReadOption{}

	_, err := r.client.Read("Sheet1", readOption, &result)
	if err != nil {
		return nil, err
	}

	for _, row := range result {
		if row.Ktp == ktp {
			filtered = append(filtered, row)
		}
	}

	return filtered, nil
}

func (r *SteinHqLoanRepository) FindLoanById(id string) (*loan.Loan, error) {
	var result []*loan.Loan

	readOption := stein.ReadOption{}

	_, err := r.client.Read("Sheet1", readOption, &result)
	if err != nil {
		return nil, err
	}

	for _, row := range result {
		if row.ID == id {
			return row, nil
		}
	}

	return nil, loan.NewLoanNotFoundException()
}
