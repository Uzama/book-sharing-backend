package config

type App struct {
	Port int    `yaml:"service-port"`
	Host string `yaml:"service-host"`
}

func (app *App) Parse() error {

	return nil
}
