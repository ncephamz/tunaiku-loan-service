package loan

type LoanNotFoundException struct{}

func NewLoanNotFoundException() *LoanNotFoundException {
	return &LoanNotFoundException{}
}

func (e *LoanNotFoundException) Error() string {
	return "Loan not found."
}
