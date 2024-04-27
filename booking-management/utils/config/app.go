package config

type App struct {
	Port int    
	Host string
}

func (app *App) Parse() error {

	port, err := getEnvInt("PORT")
	if err != nil {
		return err
	}

	host, err := getEnvString("HOST")
	if err != nil {
		return err
	}

	app.Port = port
	app.Host = host

	return nil
}
