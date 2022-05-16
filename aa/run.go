package aa

type Job interface {
	Run(*App) error
}

func (app *App) Run(jobs ...Job) {
	for _, serve := range jobs {
		if err := serve.Run(app); err != nil {
			panic(err)
		}
	}
}
