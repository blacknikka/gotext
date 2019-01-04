package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

const (
	left = iota
	center
	right
)

var dataL []int
var dataC []int
var dataR []int
var result []int
var cont bool

var deleteListL []int
var deleteListC []int
var deleteListR []int

func init() {
	result = []int{}
	cont = false
	deleteListL = []int{}
	deleteListC = []int{}
	deleteListR = []int{}
}

// AnalyzeExec 解析実行する
func AnalyzeExec() {
	// 解析を行う

	// idx := 0
	for i := 0; i < len(Inst().Table.Table[0]); i++ {
		for j := 0; j < len(Inst().Table.Table[1]); j++ {
			for k := 0; k < len(Inst().Table.Table[2]); k++ {
				fmt.Println("i:" + strconv.Itoa(i))
				fmt.Println("j:" + strconv.Itoa(j))
				fmt.Println("k:" + strconv.Itoa(k))
				if analyze(i, j, k) == false {
					// このパターンは異常として記録
				}
				fmt.Println(Inst().Result[len(Inst().Result)-1])
			}
		}
	}

	// // 残り情報
	// fmt.Println(dataL)
	// fmt.Println(dataC)
	// fmt.Println(dataR)
}

func analyze(ofsL int, ofsC int, ofsR int) bool {
	// 領域を作る
	dataL = make([]int, len(Inst().Table.Table[0]))
	dataC = make([]int, len(Inst().Table.Table[1]))
	dataR = make([]int, len(Inst().Table.Table[2]))
	copy(dataL, Inst().Table.Table[0])
	copy(dataC, Inst().Table.Table[1])
	copy(dataR, Inst().Table.Table[2])

	dataL = makeTargetArray(dataL, ofsL)
	dataC = makeTargetArray(dataC, ofsC)
	dataR = makeTargetArray(dataR, ofsR)
	ofsLP0 := 0
	ofsLP1 := 1
	ofsLP2 := 2
	ofsCP0 := 0
	ofsCP1 := 1
	ofsCP2 := 2
	ofsRP0 := 0
	ofsRP1 := 1
	ofsRP2 := 2

	ret := true

	result = make([]int, 0)

	// combo counter
	combo := 0

	for true {
		cont = false
		deleteListL = []int{}
		deleteListC = []int{}
		deleteListR = []int{}

		// L
		if dataL[ofsLP0] == dataL[ofsLP1] {
			pushResult(dataL[ofsLP0], left, ofsLP0, left, ofsLP1)
		}
		if dataL[ofsLP1] == dataL[ofsLP2] {
			pushResult(dataL[ofsLP1], left, ofsLP1, left, ofsLP2)
		}

		// C
		if dataC[ofsCP0] == dataC[ofsCP1] {
			pushResult(dataC[ofsCP0], center, ofsCP0, center, ofsCP1)
		}
		if dataC[ofsCP1] == dataC[ofsCP2] {
			pushResult(dataC[ofsCP1], center, ofsCP1, center, ofsCP2)
		}

		// R
		if dataR[ofsRP0] == dataR[ofsRP1] {
			pushResult(dataR[ofsRP0], right, ofsRP0, right, ofsRP1)
		}
		if dataR[ofsRP1] == dataR[ofsRP2] {
			pushResult(dataR[ofsRP1], right, ofsRP1, right, ofsRP2)
		}

		// 横
		if dataL[ofsLP0] == dataC[ofsCP0] {
			pushResult(dataL[ofsLP0], left, ofsLP0, center, ofsCP0)
		}
		if dataL[ofsLP1] == dataC[ofsCP1] {
			pushResult(dataL[ofsLP1], left, ofsLP1, center, ofsCP1)
		}
		if dataL[ofsLP2] == dataC[ofsCP2] {
			pushResult(dataL[ofsLP2], left, ofsLP2, center, ofsCP2)
		}

		if dataR[ofsRP0] == dataC[ofsCP0] {
			pushResult(dataR[ofsRP0], right, ofsRP0, center, ofsCP0)
		}
		if dataR[ofsRP1] == dataC[ofsCP1] {
			pushResult(dataR[ofsRP1], right, ofsRP1, center, ofsCP1)
		}
		if dataR[ofsRP2] == dataC[ofsCP2] {
			pushResult(dataR[ofsRP2], right, ofsRP2, center, ofsCP2)
		}

		if cont == false {
			// 条件成立しなかったらループを抜ける
			break
		} else {
			combo++
			// 条件成立の場合には要素を消して繰り返し
			if removeItem() == false {
				// 異常として終了させる
				ret = false
				break
			}
		}
	}

	// 結果の記録
	Inst().Result = append(Inst().Result, ResultItem{
		Max:      combo,
		GetItems: result,
		IsError:  ret != false,
	})

	return ret
}

func makeTargetArray(array []int, ofs int) []int {
	ret := array[ofs:]
	ret = append(ret[:], array[:ofs]...)
	return ret
}

func pushResult(get int, pos1 int, pos1ofs int, pos2 int, pos2ofs int) {
	// 繰り返しのフラグを立てる
	cont = true
	result = append(result[:], get)

	// 一致した要素のリストを作る
	var list1, list2 *[]int
	switch pos1 {
	case left:
		list1 = &deleteListL
	case center:
		list1 = &deleteListC
	case right:
		list1 = &deleteListR
	}

	switch pos2 {
	case left:
		list2 = &deleteListL
	case center:
		list2 = &deleteListC
	case right:
		list2 = &deleteListR
	}

	*list1 = append(*list1, pos1ofs)
	*list2 = append(*list2, pos2ofs)
}

// 条件に一致した要素を削除する処理
func removeItem() bool {
	// 重複を削除
	deleteListL = removeDistinct(deleteListL)
	deleteListC = removeDistinct(deleteListC)
	deleteListR = removeDistinct(deleteListR)

	// ソート
	sort.Sort(sort.IntSlice(deleteListL))
	sort.Sort(sort.IntSlice(deleteListC))
	sort.Sort(sort.IntSlice(deleteListR))

	// 削除するリストができたので、要素を元データから削除する
	removeFromArray()

	ret := true
	if len(dataL) < 3 || len(dataC) < 3 || len(dataR) < 3 {
		// 削除した結果要素が３つより小さくなるとエラーとする
		ret = false
	}

	return ret
}

// 重複を削除する処理
func removeDistinct(array []int) []int {
	arr := make(map[int]bool)
	ret := []int{}
	for _, v := range array {
		if !arr[v] {
			arr[v] = true
			ret = append(ret, v)
		}
	}

	return ret
}

func removeFromArray() {
	for i := (len(deleteListL) - 1); i >= 0; i-- {
		// 配列から削除
		dataL = unset(dataL, deleteListL[i])
	}

	for i := (len(deleteListC) - 1); i >= 0; i-- {
		// 配列から削除
		dataC = unset(dataC, deleteListC[i])
	}

	for i := (len(deleteListR) - 1); i >= 0; i-- {
		// 配列から削除
		dataR = unset(dataR, deleteListR[i])
	}
}

// 配列を削除する
func unset(s []int, i int) []int {
	if i >= len(s) {
		log.Fatal("削除エラー")
		return s
	}

	ret := []int{}
	ret = append([]int{s[len(s)-1]}, s[:i]...)
	ret = append(ret, s[i+1:len(s)-1]...)

	return ret
}
