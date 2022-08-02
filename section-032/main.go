package main

// インターフェースの注意点

import "fmt"

type MyErr struct{}

func (me *MyErr) Error() string { return "" }

// Apply関数は常にnilが代入されたerr変数を返す関数です
// しかし返り値をnilと比較してもfalseになる
// これは戻り値がerrorインターフェースのため、実際の
// 戻り値は「型が*MyErr型で値がnilのerrorインターフェースを満たす値」が返されるため
func Apply() error {
	var err *MyErr = nil
	//戻り値に型情報が含まれているので予想外の挙動をする
	return err
}

func Apply2() error {
	var err error = nil
	// 結果的には問題ないが、明示的にnilをreturnしてミスを防ぐほうが良い
	return err
}

func main() {
	fmt.Println(Apply() == nil)  // false
	fmt.Println(Apply2() == nil) // true
}
