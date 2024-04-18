package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RunBlockShadowingControl() {
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

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// for文の選択
	// 基本的にはfor range
	// 合成型の処理で全要素のいてレーションをしない場合
	// for版
	evenVals = []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		// iが0ならばスキップして次のループへ
		if i == 0 {
			continue
		}
		fmt.Println(i, v)
		// iが最後まで回ったら抜ける
		if i == len(evenVals) {
			break
		}
	}

	evenVals = []int{2, 4, 6, 8, 10, 12}
	// iはから（0をスキップと同義）, iは配列スライスの分だけループ
	for i := 1; i < len(evenVals); i++ {
		fmt.Println(i, evenVals[i])
	}

	// switch
	// Goのswitchはフォールスルーしない
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	switch size := r.Intn(10); size {
	case 1, 2, 3, 4:
		fmt.Printf("%dは1から4である", size)
	case 5, 6, 7, 8:
		fmt.Printf("%dは5から8である", size)
	default:
		fmt.Printf("%dは9から10である", size)
	}

	// ブランクswitch
	words := []string{"hi", "hello", "world"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println("短い単語")
		case wordLen > 10:
			fmt.Println("長い単語")
		default:
			fmt.Println("ちょうど良い単語")
		}
	}
}
