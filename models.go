package main

type borrower struct {
	Id                     string                  `json:"id"`
	RegistratonID          string                  `json:"registrationId"`
	Name                   string                  `json:"name"`
	Email                  string                  `json:"email"`
	FinancialReports       []financialReport       `json:"financialReport"`
	ComplianceCertificates []complianceCertificate `json:"complianceCertifiacte"`
	Requests               []request               `json:"Requests"`
	ReserveReports         []reserveReport         `json:"reserveReports"`
	LoanPacks              []loanPackage           `json:"loanPacks"`
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
	Requests       []request       `json:"requests"`
	ReserveReports []reserveReport `json:"reserveReports"`
}
type request struct {
	Id         string `json:"id"`
	BorrowerId string `json:"borrowerId"`
	Status     string `json:"status"`
}

type reserveReport struct {
	RequestId        string  `json:"requestId"`
	Id               string  `json:"id"`
	Date             string  `json:"date"`
	EngineerId       string  `json:"engineerId"`
	BorrowerId       string  `json:"borrowerId"`
	DevelopedCrude   float64 `json:"developedCrude"`
	UndevelopedCrude float64 `json:"undevelopedCrude"`
}
type document struct {
	Id      string `json:"id"`
	DocName string `json:"docName"`
}
type loanPackage struct {
	Id                    string                `json:"id"`
	BorrowerId            string                `json:"borrowerId"`
	FinancialReport       financialReport       `json:"financialReport"`
	ComplianceCertificate complianceCertificate `json:"complianceCertificate"`
	ReserveReport         reserveReport         `json:"reserveReport"`
	Documents             []document            `json:"documents"`
	AmountRequested       float64               `json:"amountRequested"`
	BorrowerName          string                `json:"borrowerName"`
	Status                string                `json:"status"`
}

type auditor struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	FinancialReports []financialReport `json:"financialReport"`
}

type administrativeAgent struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	LoanPackage []loanPackage `json:"loanPackage"`
}
