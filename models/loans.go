package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Loan struct {
	LoanNumber  string  `json:"loan_number"`
	KtpNumber   string  `json:"ktp_number"`
	DebiturName string  `json:"debitur_name"`
	PhoneNumber string  `json:"phone_number"`
	ProductCode string  `json:"product_code"`
	Arrears     float32 `json:"arrears"`
	Tenor       uint32  `json:"tenor"`
	Interest    float32 `json:"interest"`
	gorm.Model
}

func (con *Loan) Validate() (map[string]interface{}, bool) {

	resp := make(map[string]interface{})

	if con.LoanNumber == "" {
		resp["status"] = false
		resp["message"] = "Loan number is required."
		return resp, false
	}

	if con.KtpNumber == "" {
		resp["status"] = false
		resp["message"] = "KTP number is required."
		return resp, false
	}

	if con.DebiturName == "" {
		resp["status"] = false
		resp["message"] = "Debitur name is required."
		return resp, false
	}

	if con.PhoneNumber == "" {
		resp["status"] = false
		resp["message"] = "Phone number is required."
		return resp, false
	}

	if con.Arrears == 0 {
		resp["status"] = false
		resp["message"] = "Arrears is required."
		return resp, false
	}

	if con.Tenor == 0 {
		resp["status"] = false
		resp["message"] = "Tenor is required."
		return resp, false
	}

	if con.Interest == 0 {
		resp["status"] = false
		resp["message"] = "Interest is required."
		return resp, false
	}

	resp["status"] = true
	resp["message"] = "All required fields(s) present"
	return resp, true
}

func (loan *Loan) Create() map[string]interface{} {

	resp := make(map[string]interface{})
	con := &Loan{}
	GetConn().Where("loan_number = ?", loan.LoanNumber).First(con)
	if con.LoanNumber != "" {
		resp["status"] = false
		resp["message"] = "This loan number already exists"
		return resp
	}

	GetConn().Create(loan)
	resp["status"] = true
	resp["message"] = "Loan has been created."
	resp["loan"] = GetLoan(loan.ID)
	return resp
}

func (loan *Loan) Update(data *Loan) map[string]interface{} {

	resp := make(map[string]interface{})

	GetConn().Model(&loan).Updates(data)
	resp["status"] = true
	resp["message"] = "Loan has been updated."
	resp["loan"] = GetLoan(loan.ID)
	fmt.Println(data)
	return resp
}

func (loan *Loan) Delete(data *Loan) map[string]interface{} {

	resp := make(map[string]interface{})

	GetConn().Where("id = ?", data.ID).Delete(&loan)
	resp["status"] = true
	resp["message"] = "Loan has been deleted."
	resp["loan"] = GetLoan(loan.ID)
	return resp
}
