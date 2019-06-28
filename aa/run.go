package aa

type Job interface {
	Run(*Aa) error
}

func (app *Aa) Run(jobs ...Job) {
	for _, serve := range jobs {
		if err := serve.Run(app); err != nil {
			panic(err)
		}
	}
}
