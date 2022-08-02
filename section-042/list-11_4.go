package main

// 実装の詳細
type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error { return nil }

// 上位階層が定義する抽象
type OrderService interface {
	Apply(int) error
}

// 上位階層の利用者側の型
type Application struct {
	os OrderService
}

func (app *Application) Apply(os OrderService, id int) error {
	return os.Apply(id)
}

// メソッド（関数）呼び出し時にDIする方法
func main() {
	app := &Application{}
	app.Apply(&ServiceImpl{}, 19)
}
