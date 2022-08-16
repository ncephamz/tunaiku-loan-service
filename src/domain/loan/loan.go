package loan

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Loan struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Ktp               string `json:"ktp"`
	LoanAmount        string `json:"loan_amount"`
	LoanPeriodInMonth string `json:"loan_period_in_month"`
	LoanPurpose       string `json:"loan_purpose"`
	DateOfBirth       string `json:"date_of_birth"`
	Sex               string `json:"sex"`
}

func NewLoan(loan *Loan) []Loan {
	loans := []Loan{
		{
			ID:                uuid.New().String(),
			Name:              loan.Name,
			Ktp:               loan.Ktp,
			LoanAmount:        loan.LoanAmount,
			LoanPeriodInMonth: loan.LoanPeriodInMonth,
			LoanPurpose:       loan.LoanPurpose,
			DateOfBirth:       loan.DateOfBirth,
			Sex:               loan.Sex,
		},
	}
	return loans
}

func (l *Loan) Validate() error {
	purposes := "vacation renovation electronics wedding rent car investment"

	if l.Name == "" {
		return errors.New("Required Name")
	}

	if len(strings.Split(l.Name, " ")) < 2 {
		return errors.New("Name field should include at least two names(first and last name)")
	}

	ktpPatern := regexp.MustCompile("^\\d{6}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\\d{2}\\d{4}$")
	if l.Ktp == "" {
		return errors.New("Required KTP")
	}

	if ktpPatern.Match([]byte(l.Ktp)) == false {
		return errors.New("KTP should match the formula: Men's KTP number follows their date of birth: XXXXXXDDMMYYXXXX, Woman's KTP number follows similar logic, but to DD +40 is always added.")
	}

	if l.LoanAmount == "" {
		return errors.New("Required Loan Amount")
	}

	loanAmount, _ := strconv.Atoi(l.LoanAmount)
	if loanAmount < 1000000 || loanAmount > 10000000 {
		return errors.New("Loan amount between 1000000 and 10000000")
	}

	if l.LoanPurpose == "" {
		return errors.New("Required Loan Purpose")
	}

	if strings.Contains(purposes, l.LoanPurpose) == false {
		return errors.New("Purpose includes at least one of the following words: vacation, renovation, electronics, wedding, rent, car, investment.")
	}

	return nil
}
