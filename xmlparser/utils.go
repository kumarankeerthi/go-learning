package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kumarankeerthi/go-learning/xmlparser/model"
)

func FetchXMLFiles(path string, channelOfFiles chan<- os.FileInfo) {

	f, err := os.Open(path)
	CheckError(err, "Reading input Directory")
	listOfFiles, err := f.Readdir(-1)
	CheckError(err, "Reading input Directory")
	//listOfFiles, err := ioutil.ReadDir(path)
	//CheckError(err, "Reading input Directory")

	for _, file := range listOfFiles {
		fileName := file.Name()
		if "xml" == fileName[len(fileName)-3:] {
			channelOfFiles <- file
		}
	}
	close(channelOfFiles)

}

func ParseXMLFiles(path string, channelOfFiles <-chan os.FileInfo, channelOfPolicies chan<- model.Policy) {
	p := model.Policy{}
	defer close(channelOfPolicies)
	for file := range channelOfFiles {
		p = model.Policy{}
		data, err := ioutil.ReadFile(path + "/" + file.Name())
		CheckError(err, "Reading xml file")

		err = xml.Unmarshal([]byte(data), &p)
		CheckError(err, "Unmarshalling xml file")
		channelOfPolicies <- p
	}
}

func CheckError(e error, step string) {
	if e != nil {
		fmt.Println(step, e)
		panic("system failure at")
	}
}
