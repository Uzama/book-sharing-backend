package config

type Database struct {
	User               string 
	Password           string 
	Host               string 
	Port               int    
	Database           string   
}

func (db *Database) Parse() error {

	user, err := getEnvString("DB_USER")
	if err != nil {
		return err
	}

	password, err := getEnvString("DB_PASSWORD")
	if err != nil {
		return err
	}

	host, err := getEnvString("DB_HOST")
	if err != nil {
		return err
	}

	port, err := getEnvInt("DB_PORT")
	if err != nil {
		return err
	}

	name, err := getEnvString("DB_NAME")
	if err != nil {
		return err
	}

	db.User = user
	db.Password = password
	db.Host = host
	db.Port = port
	db.Database = name

	return nil
}
