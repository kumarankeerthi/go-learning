package main

import "encoding/xml"

type Document struct {
	DocumentId    int
	TransactionId string
	DocumentType  string `xml:"DocType"`
}

type Transaction struct {
	TransactionId string
	Comment       string
}

type Policy struct {
	XMLName      xml.Name      `xml:"Policy"`
	PolicyId     int           `xml:"ID,attr"`
	CaseNumber   int           `xml:"CaseNum"`
	Transactions []Transaction `xml:"Transactions>Transaction"`
	Documents    []Document    `xml:"Documents>Document"`
}
