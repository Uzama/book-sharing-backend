package config

type Database struct {
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	Database           string `yaml:"database"`
	IdleConnection     int    `yaml:"idle-connection"`
	OpenConnection     int    `yaml:"open-connection"`
	ConnectionLifeTime int    `yaml:"connection-life-time"`
	ConnectionIdleTime int    `yaml:"connection-idle-time"`
	ReadTimeout        int    `yaml:"read-timeout"`
	WriteTimeout       int    `yaml:"write-timeout"`
	Timeout            int    `yaml:"timeout"`
}

func (db *Database) Parse() error {

	return nil
}
