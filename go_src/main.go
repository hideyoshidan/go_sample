package main

import (
    "fmt"
)

func main() {

	test,test2 := sampleFunc("これは", "テストです。")
	fmt.Println(test + " " + test2)
}

func sampleFunc(sample, sample2 string) (string, string) {
	sample3 := sample + sample2
	sample4 := "成功しました。"
	return sample3, sample4
}