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
	Cases                  []Case                  `json:"Cases"`
}

type financialReport struct {
	Id           string  `json:"id"`
	RequestId    string  `json:"requestId"`
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
	Id             string           `json:"id"`
	Name           string           `json:"name"`
	RegistratonID  string           `json:"registrationId"`
	Email          string           `json:"email"`
	Requests       []reserveRequest `json:"requests"`
	ReserveReports []reserveReport  `json:"reserveReports"`
}
type request struct {
	Id         string `json:"id"`
	BorrowerId string `json:"borrowerId"`
	RequestTo  string `json:"requestTo"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Date       string `json:"date"`
}

type reserveRequest struct {
	LoanId     string `json:"loanId"`
	Id         string `json:"id"`
	BorrowerId string `json:"borrowerId"`
	EngineerId string `json:"engineerId"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Date       string `json:"date"`
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
type Case struct {
	Id                    string                `json:"id"`
	BorrowerId            string                `json:"borrowerId"`
	AdministrativeAgentId string                `json:"administrativeAgentId"`
	FinancialReports      []financialReport     `json:"financialReports"`
	ComplianceCertificate complianceCertificate `json:"complianceCertificate"`
	RequestReserveReport  request               `json:"requestReserveReport"`
	ReserveReport         reserveReport         `json:"reserveReport"`
	Documents             []document            `json:"documents"`
	AmountRequested       float64               `json:"amountRequested"`
	BorrowerName          string                `json:"borrowerName"`
	Status                string                `json:"status"`
}

type proposal struct {
	Id       string  `json:"id"`
	CaseId   string  `json:"CaseId"`
	LenderId string  `json:"lenderId"`
	Amount   float64 `json;"amount"`
	Status   string  `json:"status"`
}

type auditor struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	Requests         []request         `json:"requests"`
	FinancialReports []financialReport `json:"financialReport"`
}

type administrativeAgent struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cases []Case `json:"Cases"`
}

type lender struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Cases     []Case     `json:"Cases"`
	Proposals []proposal `json:"proposals"`
}
