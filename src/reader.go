package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Read ファイルを読み込んで、オブジェクトを取得します。
func Read(path string, masterPath string) {
	// データの読み込み
	records := fileReader(path)

	// マスターデータの読み込み
	readMaster(masterPath)

	// 各行でループ
	inst := Inst()
	for _, v := range records {
		// 読み込んだデータを登録する
		data1 := GetMasterIndex(v[0])
		data2 := GetMasterIndex(v[1])
		data3 := GetMasterIndex(v[2])
		inst.Table.Table = append(inst.Table.Table[:], [3]int{
			data1,
			data2,
			data3})
	}

	for _, v := range Inst().Table.Table {
		fmt.Println(v)
	}
}

func readMaster(masterPath string) {
	records := fileReader(masterPath)

	for idx, v := range records {
		i, _ := strconv.Atoi(v[1])

		// itemを設定する
		item := Item{Name: v[0], Score: i}
		Inst().Master.Dic[idx] = item

		// 名称からの逆引きも登録しておく
		Inst().Master.DicByName[v[0]] = idx
	}
}

func fileReader(path string) [][]string {
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

	return records
}
