package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"
)

// Read ファイルを読み込んで、オブジェクトを取得します。
func Read(path string, masterPath string) {
	// ファイル読み込み
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		// ファイルオープンエラー
		log.Fatal(err)
	}

	// CSVのReaderを用意
	r := csv.NewReader(strings.NewReader(string(contents)))

	// デリミタ(TSVなら\t, CSVなら,)設定
	r.Comma = '\t'

	// コメント設定
	r.Comment = '#'

	// 全部読みだす
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 各行でループ
	inst := Inst()
	for _, v := range records {
		// 読み込んだデータを登録する
		inst.Table.Table = append(inst.Table.Table[:], [3]Item{
			Item{Name: v[0]},
			Item{Name: v[1]},
			Item{Name: v[2]}})
	}
}
