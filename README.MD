# GoLang Rest Api Development

A simple loans manager.

## Required Packages
- Dependency management
    * [dep](https://github.com/golang/dep)
    * [migrate](https://github.com/golang-migrate/migrate)
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- ORM   
    * [gorm](https://github.com/jinzhu/gorm)
- Routing
    * [mux](https://github.com/gorilla/mux)

## Quick Run Project
First clone https://github.com/ajengmira/goloan or copy source then go to goloan folder.  

```
##go to your db editor
create database and table
to execute query lookup .\db\migrations\20210720136960_CreateLoanTable.up

##go to your console
cd loan

##go mod tidy for sync package

go mod tidy

go run main.go
```

## Docker Build
sudo docker build -t goloan

#####################################################################################

## API ENDPOINTS

### All loan
- Path : `/api/loans`
- Method: `GET`
- Response: `200`

### Loan By Id
- Path : `/api/loan/{id}`
- Method: `GET`
- Response: `200`

### Create loan
- Path : `/api/loan`
- Method: `POST`
- Fields: `loan_number, debitur_name, phone_number, product_code, arrears, tenor, interest`
- Response: `200`

### Details a loan
- Path : `/api/loan/{id}`
- Method: `GET`
- Response: `200`

### Update loan
- Path : `/api/loan/{id}`
- Method: `PUT`
- Fields: `loan_number, debitur_name, phone_number, product_code, arrears, tenor, interest`
- Response: `200`

### Approve loan
- Path : `/api/loan/approve/{id}`
- Method: `PUT`
- Fields: `loan_number`
- Response: `200`

### Delete loan
- Path : `/api/loan/{id}`
- Method: `DELETE`
- Response: `200`

