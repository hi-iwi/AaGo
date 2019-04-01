package aa

type Runner interface {
	Run(*Aa) error
}

func (app *Aa) Run(runners ...Runner) {
	for _, serve := range runners {
		if err := serve.Run(app); err != nil {
			panic(err)
		}
	}
}
