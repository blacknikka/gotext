package main

// Item 各アイテム
type Item struct {
	Name  string
	Score int
}

// Master マスター情報
type Master struct {
	Dic       map[int]Item
	DicByName map[string]int
}

// GetMasterIndex マスターから情報を取得する
func GetMasterIndex(name string) int {
	_, ok := Inst().Master.DicByName[name]

	ret := -1
	if ok == true {
		ret = Inst().Master.DicByName[name]
	}

	return ret
}
