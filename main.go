package main

import "fmt"

func main() {
	// capの増え方
	// capとlenが等しくなると2倍に増やす
	// 要素数が1024を超えると1.25倍の増加に変わる
	var x []int
	fmt.Println(len(x), cap(x)) // 0 0
	x = append(x, 10)
	fmt.Println(len(x), cap(x)) // 1 1
	x = append(x, 20)
	fmt.Println(len(x), cap(x)) // 2 2
	x = append(x, 30)
	fmt.Println(len(x), cap(x)) // 3 4
	x = append(x, 40)
	fmt.Println(len(x), cap(x)) // 4 4
	x = append(x, 50)
	fmt.Println(len(x), cap(x)) // 5 8

	// len5のスライスを定義する
	s1 := make([]int, 5)
	// appendすると最初の値に10が入ると思いきや
	s1 = append(s1, 10)
	// 実行結果はゼロ値で初期化された0の後に入る
	fmt.Println(s1, len(s1), cap(s1)) //  [0 0 0 0 0 10] 6 10

	// lenを0にして空にしておくと意図通りの挙動を実現できそう
	s2 := make([]int, 0, 10)
	s2 = append(s2, 10)
	s2 = append(s2, 20)
	s2 = append(s2, 30)
	fmt.Println(s2, len(s2), cap(s2)) // [10 20 30] 3 10

	// スライス生成方法
	// 初期値がある場合はスライスリテラルを使う
	s3 := []int{10, 20, 30}
	fmt.Println(s3, len(s3), cap(s3))

	// 必要なサイズがわかっている場合
	s4 := make([]int, 0, 10)
	fmt.Println(s4, len(s4), cap(s4))

	// 長さ0でcapを指定してmake
	// 要素追加時にappend、最初に容量を指定しないので余分なゼロ値が入らない
	s5 := make([]int, 0)
	fmt.Println(s5, len(s5), cap(s5))
	s5 = append(s5, 20)
	fmt.Println(s5, len(s5), cap(s5))
	s5 = append(s5, 30)
	fmt.Println(s5, len(s5), cap(s5))
}
