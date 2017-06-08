package main

type borrower struct {
	Id                     string                  `json:"id"`
	RegistratonID          string                  `json:"registrationId"`
	Name                   string                  `json:"name"`
	Email                  string                  `json:"email"`
	FinancialReports       []financialReport       `json:"financialReport"`
	ComplianceCertificates []complianceCertificate `json:"complianceCertifiacte"`
	ReserveReports         []reserveReport         `json:"reserveReports"`
}

type financialReport struct {
	Id           string  `json:"id"`
	BorrowerId   string  `json:"borrowerId"`
	Date         string  `json:"date"`
	LoanAmount   float64 `json:"loanAmount"`
	CreditPeriod int     `json:"creditPeriod"`
	Status       string  `json:"status"`
}

type complianceCertificate struct {
	Id         string `json:"id"`
	Date       string `json:"date"`
	BorrowerId string `json:"borrowerId"`
}

type engineer struct {
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	RegistratonID  string          `json:"registrationId"`
	Email          string          `json:"email"`
	ReserveReports []reserveReport `json:"reserveReports"`
}

type reserveReport struct {
	Id               string  `json:"id"`
	Date             string  `json:"date"`
	EngineerId       string  `json:"engineerId"`
	BorrowerId       string  `json:"borrowerId"`
	DevelopedCrude   float64 `json:"developedCrude"`
	UndevelopedCrude float64 `json:"undevelopedCrude"`
}

type loanPackage struct {
	BorrowerId            string                `json:"borrowerId"`
	FinancialReport       financialReport       `json:"financialReport"`
	ComplianceCertificate complianceCertificate `json:"complianceCertificate"`
	ReserveReport         reserveReport         `json:"reserveReport"`
	AmountRequested       string                `json:"amountRequested"`
	BorrowerName          string                `json:"borrowerName"`
	Status                string                `json:"status"`
}

type auditor struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	FinancialReports []financialReport `json:"financialReport"`
}
