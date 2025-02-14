package app

type Sayer interface {
	Say(something string)
}

type App struct {
	sayer Sayer
}

func New(sayer Sayer) *App {
	return &App{
		sayer: sayer,
	}
}

func (a *App) Run(v string) {
	a.sayer.Say("i say: " + v)
}
