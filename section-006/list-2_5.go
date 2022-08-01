package main

import (
	"context"
	"fmt"
	"time"
)

// キャンセルされるまで待ちたい
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go func() { fmt.Println("別ゴルーチン") }()
	fmt.Println("STOP")
	<-ctx.Done() // キャンセルされるまで待つ
	fmt.Println("そして時は動き出す")
}
