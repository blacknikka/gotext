package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	if len(args) != 2 {
		// 実行エラー
		log.Fatal("引数エラー")
		return
	}

	// read file
	Read(args[0], args[1])

	// 解析
	AnalyzeExec()
}
