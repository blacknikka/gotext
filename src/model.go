package main

// Item 各アイテム
type Item struct {
	Name string
}

// ItemMaster 各アイテムのマスター
type ItemMaster struct {
	Items []Item
}

// TableModel 盤面
type TableModel struct {
	Table [][3]Item
}

// ResultItem 結果の要素
type ResultItem struct {
	Max      int
	GetItems []Item
}

// Result 結果
type Result struct {
	Result []ResultItem
}

// InstType インスタンス
type InstType struct {
	Table  TableModel
	Result Result
}

var inst *InstType = &InstType{}

// Inst シングルトンインスタンス返却
func Inst() *InstType {
	return inst
}
