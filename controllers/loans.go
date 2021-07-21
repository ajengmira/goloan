package controllers

import (
	"net/http"
	"goloan/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
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
	}else {
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
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "loan" : loan}
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
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "loan" : loan}
	 
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
	
	loan := new(models.Loan)
	loan.LoanNumber = loanNumber
	loan.KtpNumber = ktpNumber
	loan.DebiturName = debiturName
	loan.PhoneNumber = phoneNumber
	loan.ProductCode = productCode
	loan.Arrears = arrears
	loan.Tenor = tenor
	loan.Interest = interest
	loans := mdl.Update(loan)
	
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "loans" : loans}
	var jsonData, err = json.Marshal(resp)
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
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "loans" : loans}
	var jsonData, err = json.Marshal(resp)
	if err != nil {
	    fmt.Println(err.Error())
	    return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}