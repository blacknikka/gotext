package main

type Item struct {
	Name string
}

// 各アイテムのマスター
type ItemMaster struct {
	Items []Item
}

// 盤面
type TableModel struct {
	Table [][3]Item
}

// 結果の要素
type ResultItem struct {
	Max      int
	GetItems []Item
}

// 結果
type Result struct {
	Result []ResultItem
}

// inst
type InstType struct {
	Table  TableModel
	Result Result
}

var inst *InstType = &InstType{}

func Inst() *InstType {
	return inst
}
