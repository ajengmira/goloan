package controllers

import (
	"encoding/json"
	"fmt"
	"goloan/app/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateLoan = func(w http.ResponseWriter, r *http.Request) {

	loan := &models.Loan{}
	json.NewDecoder(r.Body).Decode(loan)

	if resp, ok := loan.Validate(); !ok {
		var jsonData, err = json.Marshal(resp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		resp = loan.Create()
		var jsonData, err = json.Marshal(resp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

var GetLoans = func(w http.ResponseWriter, r *http.Request) {

	loan := models.GetLoans()
	resp := map[string]interface{}{"status": true, "message": "Request success", "loan": loan}
	var jsonData, err = json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

var GetLoan = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	loan := models.GetLoan(uint(id))
	resp := map[string]interface{}{"status": true, "message": "Request success", "loan": loan}

	var jsonData, err = json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

var UpdateLoan = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	loanNumber := r.FormValue("loan_number")
	ktpNumber := r.FormValue("ktp_number")
	debiturName := r.FormValue("debitur_name")
	phoneNumber := r.FormValue("phone_number")
	productCode := r.FormValue("product_code")
	arrears := r.FormValue("arrears")
	tenor := r.FormValue("tenor")
	interest := r.FormValue("interest")

	// get the param id
	mdl := models.GetLoan(uint(id))

	arrearsInt, err := strconv.ParseFloat(arrears, 32)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	arrearsInt_ := float32(arrearsInt)

	tenorInt, err := strconv.ParseUint(tenor, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tenorInt_ := uint32(tenorInt)

	interestInt, err := strconv.ParseFloat(interest, 32)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	interestInt_ := float32(interestInt)

	loan := new(models.Loan)
	loan.LoanNumber = loanNumber
	loan.KtpNumber = ktpNumber
	loan.DebiturName = debiturName
	loan.PhoneNumber = phoneNumber
	loan.ProductCode = productCode
	loan.Arrears = arrearsInt_
	loan.Tenor = tenorInt_
	loan.Interest = interestInt_
	loans := mdl.Update(loan)

	resp := map[string]interface{}{"status": true, "message": "Request success", "loans": loans}
	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

var DeleteLoan = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	// get the param id
	mdl := models.GetLoan(uint(id))

	loan := new(models.Loan)
	loan.ID = uint(id)
	loans := mdl.Delete(loan)
	resp := map[string]interface{}{"status": true, "message": "Request success", "loans": loans}
	var jsonData, err = json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
