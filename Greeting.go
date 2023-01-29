package main

import (
	"fmt"
)

// 配列の作り方
// 変数名 := [要素数] データ型 {データ１, データ２}
//
// 要素数を省略する場合・・・
// 変数名 := [...] データ型 {データ１, データ２}
//
// 多次元配列
// a := [2][2]string{{"aa", "aaaa"}, {"aa", "aaaa"}}

// if文 (簡易文)
// 簡易文の変数はif文内でしか使えない！
// if age := 0; age >= 20 {
// 	fmt.Println("adult")
// } else {
// 	fmt.Println("child")
// }

// 構造体
// type Student struct {
// 	name string
// 	age int
// }
//
// 代入方法
// s := Student{name: "maru", age: 26}

type Student struct {
	name string
	age int
}

func main() {
	s := Student{name: "maru", age: 26}
	fmt.Println(s)
}
