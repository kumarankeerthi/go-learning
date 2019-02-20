package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type consoFile os.File

func main1() {

	inputFileName := "C:\\Users\\kkumaran\\Desktop\\go\\confirms.txt"
	outputFolder := "C:\\Users\\kkumaran\\Desktop\\go\\"
	content, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Println("failed to open inputFileName")
		return
	}
	for i := 0; i < 3000; i++ {
		outFile, err := os.OpenFile(outputFolder+"confirms"+fmt.Sprint(i)+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("something wrong")
		}
		_, err1 := outFile.Write(content)
		if err1 != nil {
			fmt.Println("something wrong")
		}
		outFile.Close()
	}

}
