package main

func CustomApply(id int) error { return nil }

type Applicatinon struct {
	Apply func(int) error
}

func (app *Applicatinon) Run(id int) error {
	return app.Apply(id)
}

// 関数を利用したDI
func main() {
	app := &Applicatinon{Apply: CustomApply}
	app.Run(19)
}
