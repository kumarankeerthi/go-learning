package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kumarankeerthi/go-learning/xmlparser/model"
)

func PrintData(outputFilePath string, data <-chan model.Policy) {
	s := ""
	polFile, err := os.Create(outputFilePath + "\\PolicyData.txt")
	CheckError(err, "creating policy data file")

	txnFile, err := os.Create(outputFilePath + "\\TxnData.txt")
	CheckError(err, "creating policy data file")

	docFile, err := os.Create(outputFilePath + "\\DocData.txt")
	CheckError(err, "creating policy data file")

	for p := range data {
		for _, ca := range p.Cases {
			s = fmt.Sprintf("  %s ¿ %s ¿ %d", p.PolicyId, ca.CaseId, ca.CaseNumber)
			fmt.Fprintln(polFile, s)

			for _, txn := range ca.Transactions {
				s = fmt.Sprintf("  %s ¿ %d ¿ %s", p.PolicyId, ca.CaseNumber, txn.TransactionId)
				fmt.Fprintln(txnFile, s)
			}
			for _, doc := range ca.Documents {
				s = fmt.Sprintf("  %s ¿ %d ¿ %d", p.PolicyId, ca.CaseNumber, doc.DocumentId)
				fmt.Fprintln(docFile, s)
			}
		}
	}

	err = polFile.Close()
	err = txnFile.Close()
	err = docFile.Close()
}

func main() {
	inputFilePath := flag.String("inputFilePath", "C:\\Users\\kkumaran\\Desktop\\go-data\\input", "XML file location")
	outputFilePath := flag.String("outputFilePath", "C:\\Users\\kkumaran\\Desktop\\go-data\\output", "Output file location")
	flag.Parse()

	start := time.Now()

	channelOfFiles := make(chan os.FileInfo)
	channelOfPolices := make(chan model.Policy)

	go FetchXMLFiles(*inputFilePath, channelOfFiles)

	go ParseXMLFiles(*inputFilePath, channelOfFiles, channelOfPolices)

	PrintData(*outputFilePath, channelOfPolices)

	elapsed := time.Since(start)
	fmt.Printf("Processing took %s", elapsed)
}
