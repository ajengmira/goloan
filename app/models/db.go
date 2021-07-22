package models

func GetLoan(Id uint) *Loan {

	loan := &Loan{}
	GetConn().First(loan, Id)
	return loan
}

func GetLoans() *Loan {

	loan := &Loan{}
	GetConn().Order("id desc").Find(&loan)
	return loan
}
