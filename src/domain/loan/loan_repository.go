package loan

type LoanRepository interface {
	CreateLoan(loan *Loan) (*Loan, error)
	FindAllByKtp(ktp string) ([]*Loan, error)
	FindLoanById(id string) (*Loan, error)
}
