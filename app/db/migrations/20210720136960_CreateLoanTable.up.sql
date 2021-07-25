CREATE DATABASE goloan;

DROP TABLE IF EXISTS loans;

CREATE TABLE loans (
  id            INT UNSIGNED	NOT NULL AUTO_INCREMENT PRIMARY KEY,
  loan_number   VARCHAR(150)  NOT NULL,
  ktp_number  	VARCHAR(150)	NOT NULL,
  debitur_name  VARCHAR(150)	NOT NULL,
  phone_number  VARCHAR(150)	NOT NULL,
  product_code  VARCHAR(150),
  arrears  		DOUBLE	NOT NULL,
  tenor  		INT 			NOT NULL,
  interest  	DOUBLE	NOT NULL,
  status     INT,
  created_at    DATETIME,
  updated_at    DATETIME,
  deleted_at    DATETIME
);
