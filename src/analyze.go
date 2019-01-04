package main

import (
	"fmt"
	"sort"
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
	// 領域を作る
	dataL = make([]int, len(Inst().Table.Table[0]))
	dataC = make([]int, len(Inst().Table.Table[1]))
	dataR = make([]int, len(Inst().Table.Table[2]))
	copy(dataL, Inst().Table.Table[0])
	copy(dataC, Inst().Table.Table[1])
	copy(dataR, Inst().Table.Table[2])

	// 解析を行う
	analyze(0, 0, 0)

	// idx := 0
	for i := 0; i < len(dataL); i++ {
		for j := 0; j < len(dataC); j++ {
			for k := 0; k < len(dataR); k++ {

			}
		}
	}

	// リストを表示
	for _, v := range result {
		fmt.Println(v)
	}
}

func analyze(ofsL int, ofsC int, ofsR int) {
	for true {
		cont = false
		ofsLP1 := getOfsP1(ofsL, dataL)
		ofsLP2 := getOfsP2(ofsL, dataL)
		ofsCP1 := getOfsP1(ofsC, dataC)
		ofsCP2 := getOfsP2(ofsC, dataC)
		ofsRP1 := getOfsP1(ofsR, dataR)
		ofsRP2 := getOfsP2(ofsR, dataR)

		// L
		if dataL[ofsL] == dataL[ofsLP1] {
			pushResult(dataL[ofsL], left, ofsL, left, ofsLP1)
		}
		if dataL[ofsLP1] == dataL[ofsLP2] {
			pushResult(dataL[ofsLP1], left, ofsLP1, left, ofsLP2)
		}

		// C
		if dataC[ofsC] == dataC[ofsCP1] {
			pushResult(dataC[ofsC], center, ofsC, center, ofsCP1)
		}
		if dataC[ofsCP1] == dataC[ofsCP2] {
			pushResult(dataC[ofsCP1], center, ofsCP1, center, ofsCP2)
		}

		// R
		if dataR[ofsR] == dataR[ofsRP1] {
			pushResult(dataR[ofsR], right, ofsR, right, ofsRP1)
		}
		if dataR[ofsRP1] == dataR[ofsRP2] {
			pushResult(dataR[ofsRP1], right, ofsRP1, right, ofsRP2)
		}

		// 横
		if dataL[ofsL] == dataC[ofsC] {
			pushResult(dataL[ofsL], left, ofsL, center, ofsC)
		}
		if dataL[ofsLP1] == dataC[ofsCP1] {
			pushResult(dataL[ofsLP1], left, ofsLP1, center, ofsCP1)
		}
		if dataL[ofsLP2] == dataC[ofsCP2] {
			pushResult(dataL[ofsLP2], left, ofsLP2, center, ofsCP2)
		}

		if dataR[ofsR] == dataC[ofsC] {
			pushResult(dataR[ofsR], right, ofsR, center, ofsC)
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
			// 条件成立の場合には要素を消して繰り返し
			removeItem()
			break
		}
	}
}

func getOfsP1(ofs int, array []int) int {
	ofsP1 := ofs + 1
	if ofsP1 >= len(array) {
		ofsP1 -= len(array)
	}

	return ofsP1
}

func getOfsP2(ofs int, array []int) int {
	ofsP2 := ofs + 2
	if ofsP2 >= len(array) {
		ofsP2 -= len(array)
	}

	return ofsP2
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
func removeItem() {
	// 重複を削除
	deleteListL = removeDistinct(deleteListL)
	deleteListC = removeDistinct(deleteListC)
	deleteListR = removeDistinct(deleteListR)

	// ソート
	sort.Sort(sort.IntSlice(deleteListL))
	sort.Sort(sort.IntSlice(deleteListC))
	sort.Sort(sort.IntSlice(deleteListR))

	// 削除するリストができたので、要素を元データから削除する
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
