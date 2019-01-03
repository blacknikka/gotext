package main

// ResultItem 結果の要素
type ResultItem struct {
	Max      int
	GetItems []Item
}

// Table 取得したテーブル情報
type Table struct {
	Table [][3]int
}

// InstType インスタンス
type InstType struct {
	Table  Table
	Master Master
	Result []ResultItem
}

// InstType インスタンス
var inst *InstType

func init() {
	inst = &InstType{
		Table: Table{
			Table: [][3]int{},
		},
		Master: Master{
			Dic:       map[int]Item{},
			DicByName: map[string]int{},
		},
	}
}

// Inst シングルトンインスタンス返却
func Inst() *InstType {
	return inst
}
