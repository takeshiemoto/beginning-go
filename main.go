package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	s6 := []int{10, 20, 30}
	s7 := s6[:2]
	fmt.Println(s7) // [10 20] 0番目から2の手前（オフセット）まで

	// シャドーイング
	// Goではシャドーイングが可能だが極力避けるべき
	// 標準パッケージまでもシャドーイングされる可能性がある
	// linterで検知したい
	shadow := 10
	if shadow > 5 {
		fmt.Println(shadow) // 10
		shadow := 5
		fmt.Println(shadow) // 5
	}
	fmt.Println(shadow) // 10

	// if
	// 乱数生成
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// nにランダムな数字を定義
	// n == 0が条件
	// nはelseブロックまで有効である
	if n := r.Intn(10); n == 0 {
		fmt.Println("少し小さすぎます", n)
	} else if n > 5 {
		fmt.Println("少し大きすぎる")
	} else {
		fmt.Println("良い感じ")
	}

}
