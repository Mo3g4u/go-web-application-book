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

func (app *Application) SetService(os OrderService) {
	app.os = os
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

// 「setter」を用意しておいてDIする方法
func main() {
	app := &Application{}
	app.SetService(&ServiceImpl{})
	app.Apply(19)
}
