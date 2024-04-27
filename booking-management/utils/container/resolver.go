package container

import (
	"booking-management/externals/adapters"
	"booking-management/utils/config"
	"database/sql"
)

func Resolve(config config.Config) (Containers, error) {
	adapters, err := resolveAdapters(config)
	if err != nil {
		return Containers{}, err
	}

	repos, err := resolveRepositories(adapters.Db)
	if err != nil {
		return Containers{}, err
	}

	cont := Containers{
		Adapters:     adapters,
		Repositories: repos,
	}

	return cont, nil
}

func resolveAdapters(config config.Config) (Adapters, error) {

	mysql, err := adapters.NewDB(config.Database)
	if err != nil {
		return Adapters{}, err
	}

	adapters := Adapters{
		Db: mysql,
	}

	return adapters, nil
}

func resolveRepositories(db *sql.DB) (Repositories, error) {
	
	return Repositories{}, nil
}
