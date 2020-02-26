package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

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

func checkError(e error, step string) {
	if e != nil {
		fmt.Println(step)
		panic("system failure at")
	}
}

func FetchXMLFiles(path string, channelOfFiles chan<- os.FileInfo) {

	listOfFiles, err := ioutil.ReadDir(path)
	checkError(err, "Reading input Directory")

	for _, file := range listOfFiles {
		channelOfFiles <- file
	}
	close(channelOfFiles)

}

func ParseXMLFiles(path string, channelOfFiles <-chan os.FileInfo, channelOfPolicies chan<- Policy) {
	p := Policy{}
	defer close(channelOfPolicies)
	for file := range channelOfFiles {
		p = Policy{}
		//fmt.Println("Parsing xml file name :", file.Name())
		data, err := ioutil.ReadFile(path + "/" + file.Name())
		checkError(err, "Reading xml file")

		err = xml.Unmarshal([]byte(data), &p)
		checkError(err, "Unmarshalling xml file")
		channelOfPolicies <- p
	}
}

func PrintData(data <-chan Policy) {
	fmt.Println("Policy  CaseNumber  TxnId  DocId")
	for p := range data {
		fmt.Println(" ")
		fmt.Print(p.PolicyId, p.CaseNumber)
		for _, txn := range p.Transactions {
			fmt.Print("      ", txn.TransactionId)
		}
		for _, doc := range p.Documents {
			fmt.Print("     ", doc.DocumentId)
		}
		fmt.Println("  ")
	}
}

func main() {
	inputFilePath := flag.String("inputFilePath", "/Users/KumaranKeerthi/go/data", "XML file location")
	flag.Parse()

	channelOfFiles := make(chan os.FileInfo)
	channelOfPolices := make(chan Policy)
	go FetchXMLFiles(*inputFilePath, channelOfFiles)
	go ParseXMLFiles(*inputFilePath, channelOfFiles, channelOfPolices)
	PrintData(channelOfPolices)
}
