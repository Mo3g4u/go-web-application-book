package main

// context.Contect型の値にデータを含める

import (
	"context"
	"fmt"
)

// キーには空構造体を用いて独自型を使うのが一般的
// プリミティブな値は他パッケージとキーが衝突する可能性があるので避ける

type TraceID string

const ZeroTraceID = ""

type traceIDKey struct{}

// 専用のヘルパー関数を定義しておく
func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
}

func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
