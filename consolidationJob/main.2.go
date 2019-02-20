package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sync"
	"time"
)

func main() {
	var (
		folderName     = flag.String("folderName", "C:\\Users\\kkumaran\\Desktop\\go", "The folder that contains all the flat files to consolidate")
		outputFilename = flag.String("outputFilename", "test.txt", "Output file name")
		concurrency    = flag.Int("concurrency", 2, "no.of concurrent GO thread to process this request")
	)
	flag.Parse()
	fmt.Println("starting!!!!!!!!!!")
	listOfFileNames, err := ioutil.ReadDir(*folderName)
	if err != nil {
		fmt.Println("failed to read the directory")
		return
	}
	regex, err := regexp.Compile(".txt")
	if err != nil {
		return
	}
	startTime := time.Now()
	channelOfBytes := make(chan []byte)
	textFileNames := make(chan string)
	go func() {
		for _, name := range listOfFileNames {
			if !name.IsDir() {
				if regex.MatchString(name.Name()) {
					textFileNames <- *folderName + "\\" + name.Name()
				}
			}
		}
		close(textFileNames)
	}()
	var wg sync.WaitGroup
	wg.Add(*concurrency )

	go func() {
		wg.Wait()
		close(channelOfBytes)
	}()

	for i := 0; i < *concurrency; i++ {
		go func() {
			defer wg.Done()
			for textfile := range textFileNames {
				content, err := ioutil.ReadFile(textfile)
				if err != nil {
					fmt.Println("failed to open inputfile")
					return
				} else {
					channelOfBytes <- content
				}
			}

		}()

	}
	writeOutputFile(channelOfBytes, *outputFilename)
	endTime := time.Now()
	fmt.Printf("Duration %v ", endTime.Sub(startTime))
	
}

func writeOutputFile(channelOfBytes <-chan []byte, outputFilename string) {
	outFile, err := os.OpenFile(outputFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("failed to open outputilename")
		return
	}
	for content := range channelOfBytes {
		_, err := outFile.Write(content)
		if err != nil {
			fmt.Println("something wrong")
		}
	}
	outFile.Close()

}

























// channelOfBytes := make(chan []byte)
// go func() {
// 	fmt.Println("enterint first goroutine")
// 	for _, name := range listOfFileNames {
// 		if !name.IsDir() {
// 			if regex.MatchString(name.Name()) {
// 				content, err := ioutil.ReadFile(*folderName + "\\" + name.Name())
// 				if err != nil {
// 					// fmt.Println("failed to open inputfile")
// 					return
// 				} else {
// 					// fmt.Println("Added this file to channel" + name.Name())
// 					channelOfBytes <- content
// 				}
// 			}
// 		}
// 	}
// 	close(channelOfBytes)
// 	// wg.Done()
// }()

// go func() {

// 	fmt.Println("enterint 2nd goroutine")

// 	for content := range channelOfBytes {
// 		// fmt.Println("recived content")

// 		_, err := outFile.Write(content)
// 		if err != nil {
// 			fmt.Println("something wrong")
// 		}
// 		// fmt.Println("Appending to output file")

// 	}
// 	outFile.Close()
// 	wg.Done()

// }()
// 	wg.Wait()