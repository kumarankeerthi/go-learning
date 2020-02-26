package model

import "encoding/xml"

type Document struct {
	DocumentId    int
	TransactionId string `xml:"TransactionUniqueIdentifier"`
	DocumentType  string
}

type Transaction struct {
	TransactionId string `xml:"TransactionUniqueIdentifier"`
	PolicyNumber  string
}

type Case struct {
	CaseId       string        `xml:"CaseId,attr"`
	CaseNumber   int           `xml:"CaseNumber"`
	Transactions []Transaction `xml:"Transactions>Transaction"`
	Documents    []Document    `xml:"Documents>Document"`
}

type Policy struct {
	XMLName  xml.Name `xml:"Policy"`
	PolicyId string   `xml:"PolicyNumber,attr"`
	Cases    []Case   `xml:"Case"`
}
