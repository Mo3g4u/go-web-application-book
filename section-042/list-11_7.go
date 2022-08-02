package main

type Jedi interface {
	HasForce() bool
}

type Knight struct{}

// このままではコンパイルエラーになるので実装が不十分なことがわかる
var _ Jedi = (*Knight)(nil)
