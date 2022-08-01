package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	// 関数の実ロジックに入る前にcontext.Contextの状態を検証する
	if err := ctx.Err(); err != nil {
		fmt.Println("キャンセルされてる")
		return
	}
	fmt.Println("キャンセルされていない")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}
