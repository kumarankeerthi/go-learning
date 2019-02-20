package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main2() {
	var (
		folderName     = flag.String("folderName", "C:\\Users\\kkumaran\\Desktop\\go", "The folder that contains all the flat files to consolidate")
		outputFilename = flag.String("outputFilename", "test.txt", "Output file name")
		// concurrency    = flag.Int("concurrency", 1, "no.of concurrent GO thread to process this request")

	)
	flag.Parse()

	listOfFileNames, err := ioutil.ReadDir(*folderName)
	if err != nil {
		fmt.Println("failed to read the directory")
		return
	}
	regex, err := regexp.Compile(".txt")
	if err != nil {
		return
	}

	outFile, err := os.OpenFile(*outputFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("failed to open outputilename")
		return
	}
	for _, name := range listOfFileNames {
		if !name.IsDir() {
			if regex.MatchString(name.Name()) {
				fmt.Println("Opending file")
				fmt.Print(name.Name())

				content, err := ioutil.ReadFile(*folderName + "\\" + name.Name())
				if err != nil {
					fmt.Println("failed to open inputfile")
					return
				} else {
					_, err := outFile.Write(content)
					if err != nil {
						fmt.Println("something wrong")
					}
					fmt.Println("Appending to output file")
				}
			}
		}
	}
	outFile.Close()

}
