package main

import (
	"fmt"
)

var dataL []int
var dataC []int
var dataR []int
var result []int

// AnalyzeExec 解析実行する
func AnalyzeExec() {
	// 領域を作る
	dataL = make([]int, len(Inst().Table.Table[0]))
	dataC = make([]int, len(Inst().Table.Table[1]))
	dataR = make([]int, len(Inst().Table.Table[2]))
	copy(dataL, Inst().Table.Table[0])
	copy(dataC, Inst().Table.Table[1])
	copy(dataR, Inst().Table.Table[2])
	result = []int{}

	fmt.Println(dataL)
	fmt.Println(dataC)
	fmt.Println(dataR)

	analyze(0, 0, 0)

	// for true {

	// }

	// idx := 0
	for i := 0; i < len(dataL); i++ {
		for j := 0; j < len(dataC); j++ {
			for k := 0; k < len(dataR); k++ {

			}
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func analyze(ofsL int, ofsC int, ofsR int) {
	ofsLP1 := getOfsP1(ofsL, dataL)
	ofsLP2 := getOfsP2(ofsL, dataL)
	ofsCP1 := getOfsP1(ofsC, dataC)
	ofsCP2 := getOfsP2(ofsC, dataC)
	ofsRP1 := getOfsP1(ofsR, dataR)
	ofsRP2 := getOfsP2(ofsR, dataR)

	// L
	if dataL[ofsL] == dataL[ofsLP1] {
		pushResult(dataL[ofsL])
	}
	if dataL[ofsLP1] == dataL[ofsLP2] {
		pushResult(dataL[ofsLP1])
	}

	// C
	if dataC[ofsC] == dataC[ofsCP1] {
		pushResult(dataC[ofsC])
	}
	if dataC[ofsCP1] == dataC[ofsCP2] {
		pushResult(dataC[ofsCP1])
	}

	// R
	if dataR[ofsR] == dataR[ofsRP1] {
		pushResult(dataR[ofsR])
	}
	if dataR[ofsRP1] == dataR[ofsRP2] {
		pushResult(dataR[ofsRP1])
	}

	// 横
	if dataL[ofsL] == dataC[ofsC] {
		pushResult(dataL[ofsL])
	}
	if dataL[ofsLP1] == dataC[ofsCP1] {
		pushResult(dataL[ofsLP1])
	}
	if dataL[ofsLP2] == dataC[ofsCP2] {
		pushResult(dataL[ofsLP2])
	}

	if dataR[ofsR] == dataC[ofsC] {
		pushResult(dataR[ofsR])
	}
	if dataR[ofsRP1] == dataC[ofsCP1] {
		pushResult(dataR[ofsRP1])
	}
	if dataR[ofsRP2] == dataC[ofsCP2] {
		pushResult(dataR[ofsRP2])
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

func pushResult(get int) {
	result = append(result[:], get)
}
