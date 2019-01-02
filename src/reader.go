package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func Read() {
	// テスト用文字列
	str := "test\tテスト\nHello\tこんにちは"

	// CSVのReaderを用意
	r := csv.NewReader(strings.NewReader(str))

	// デリミタ(TSVなら\t, CSVなら,)設定
	r.Comma = '\t'

	// コメント設定(なんとコメント文字を指定できる!)
	r.Comment = '#'

	// 全部読みだす
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 各行でループ
	for _, v := range records {
		// 1列目
		fmt.Print(v[0])

		fmt.Print(" | ")

		// 2列目
		fmt.Println(v[1])
	}
}
