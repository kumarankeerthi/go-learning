package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "C:\\Users\\kkumaran\\Desktop"
	f, err := os.Open(path + "\\Feb bill.pdf")
	if err != nil {
		fmt.Println("error openting file")
	}
	f1, err := os.Create(path + "\\file1.pdf")
	if err != nil {
		fmt.Println("error openting file")
	}
	f2, err := os.Create(path + "\\file2.pdf")
	if err != nil {
		fmt.Println("error openting file")
	}
	f3, err := os.Create(path + "\\file3.pdf")
	if err != nil {
		fmt.Println("error openting file")
	}
	r := bufio.NewReader(f)
	w1 := bufio.NewWriter(f1)
	w2 := bufio.NewWriter(f2)

	w3 := bufio.NewWriter(f3)

	w4 := io.MultiWriter(w1, w2, w3)
	writter, err := io.Copy(w4, r)
	if err != nil {
		fmt.Println("error writing file")
	}
	fmt.Println(writter)

}
